#!/bin/bash
# set -xe

echo "Enter the Bucket name:"
read BUCKET_NAME

echo "Enter your AWS profile:"
read AWS_PROFILE

echo "Enter text to filter:"
read STRING_TO_FILTER

DOWNLOAD_DIR="/tmp/s3_download"
SORTED_DIR="/tmp/s3_sorted_files"

VALIDATE_CREDS () {
    aws s3 ls "s3://$BUCKET_NAME" --profile "$AWS_PROFILE" &>/dev/null
    if [ $? -ne 0 ]; then
      echo -e "\n   Validation Error, Please check!!!"
      echo -e "   Whether you have entered the correct AWS Profile ($AWS_PROFILE) or S3 bucket name ($BUCKET_NAME)\n"
      exit 1
    fi
}

FILTER_FILES () {
    mkdir -p "$DOWNLOAD_DIR"
    mkdir -p "$SORTED_DIR"
    echo -e "\n Filtering .txt files..."
    
    aws s3 sync s3://$BUCKET_NAME "$DOWNLOAD_DIR" --exclude "*" --include "*.txt" --profile "$AWS_PROFILE" &>/dev/null
    find "$DOWNLOAD_DIR" -name '*.txt' | while read filepath; do
        if grep -q "$STRING_TO_FILTER" "$filepath"; then
            echo -e "\033[32m Match found in file: $(basename "$filepath").\033[0m"
            sortedpath="$SORTED_DIR/$(basename "$filepath")"
            cp "$filepath" "$sortedpath"
        else
            echo -e "\033[31m No match found in file: $(basename "$filepath").\033[0m"
        fi
    done
    echo -e "\n Files containing the \"$STRING_TO_FILTER\" has been moved to $SORTED_DIR:"
}

CLEANUP_FILES () {
    echo -e "\n Starting cleanup process..."
    rm -rf "$DOWNLOAD_DIR"
    echo " Cleanup Completed!!!"
}

MAIN_FUN () {
    VALIDATE_CREDS
    FILTER_FILES
    CLEANUP_FILES
}

MAIN_FUN
