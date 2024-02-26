StorX API request test codebase

Attention:
This code is just a written codebase for future changes. It's jsut an example of how it can be implemented.

How to use it:
1. Make the .env file and fill it like in .env.example with 
    ACCESS_KEY
    SECRET_KEY  
    ENDPOINT
2. Modify main.go file in the directory for your needs.
For example:

Delete. Set bucketName and objectKey.

Get. Set bucketName and objectKey.

List buckets. Nothing to set.

List objects. Set bucketName.

Put. Set bucketName and objectKey.
Then using generated link make a request via curl:
curl -v --upload-file my-file.txt "link"

3. Using main.go from different folders make requests.