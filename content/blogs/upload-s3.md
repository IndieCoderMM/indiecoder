---
title: Upload Files to AWS S3 from React
description: Learn how to upload images/videos directly from a React Client to an AWS S3 bucket.
date: 2024-01-25
tags: [react, frontend, webdev, js]
---

In this blog, we’ll walk through how to upload images and videos (any files) directly from a React client to an AWS S3 bucket. The main benefit of this approach is reducing the server load since we don't need to upload files to our server. We only need an endpoint to get direct upload access to S3.

In this app, we have 3 parts to set up:

1. **Frontend client** with a file input field where users can upload files
2. **Server endpoint** to get signed-URL for giving access to S3 bucket
3. AWS S3 bucket where we can store our files.

## Backend

You'll need the following credentials from your AWS S3 Bucket to access the bucket from your server.

```javascript
const REGION = process.env.AWS_REGION;
const BUCKET_NAME = process.env.AWS_BUCKET_NAME;
const ACCESS_KEY_ID = process.env.AWS_ACCESS_KEY_ID;
const SECRET_ACCESS_KEY = process.env.AWS_SECRET_ACCESS_KEY;
```

### Generating a signed URL

**A signed URL** will allow the client to upload the file directly to S3 without involving your server in the file transfer. To generate a signed URL, we'll use `s3-request-presigner` from `@aws-sdk`.

This is the function for generating signed-url for S3.

```javascript
import { S3Client, PutObjectCommand } from '@aws-sdk/client-s3';
import { getSignedUrl } from '@aws-sdk/s3-request-presigner';

const s3 = new S3Client({
  region: REGION,
  credentials: {
    accessKeyId: ACCESS_KEY_ID,
    secretAccessKey: SECRET_ACCESS_KEY,
  },
});

export const generateSignedUrl = async (filename: string) => {
  const command = new PutObjectCommand({
    Bucket: BUCKET_NAME,
    Key: filename
  });

  const url = await getSignedUrl(s3, command, {
    expiresIn: 60 * 5, // 5 minutes
  });

  return url;
};
```

You can also create folders in S3 by adding the folder name in the `Key` field. For example, `Key: 'images/${filename}'`.


### Server endpoint

We'll use the above function to return the URL as the response. Here's a simple Express server endpoint.

```javascript
const express = require('express')
const app = express()

app.get('/get-signed-url',  async (req, res) => {
  const { filename } = req.query;

  const url = await generateSignedUrl(filename);

  res.status(200).send({ url });
};
```

## Frontend

For simplicity, let’s build a basic file input field using Material UI (MUI). You can use any UI library or plain HTML input field. This component will handle file selection and trigger the upload.

### FileInput Component

Here's the base structure of the `FileInput` component. 

```javascript
import { useState } from 'react';
import InsertPhotoOutlinedIcon from '@mui/icons-material/InsertPhotoOutlined';
import { MuiFileInput } from 'mui-file-input';
import { Button, InputAdornment } from '@mui/material';

const API_BASE_URL = 'http://localhost:3000';

const FileInput = ({
  label,
  placeholder,
  fileType = 'image',
}) => {
  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [success, setSuccess] = useState(false);

  const handleFileChange = (newFile) => {
    if (newFile === null) {
      return;
    }
    setFile(newFile);
    setSuccess(false);
  };

  // we'll put our upload logic in this function
  const handleUpload = async () => {
    ...
  };

  return (
      <MuiFileInput
        value={file}
        onChange={handleFileChange}
        margin="normal"
        sx={{ flex: 1 }}
        label={label ?? 'File'}
        placeholder={placeholder ?? 'Select file'}
        InputProps={{
          inputProps: {
            accept: `${fileType}/*`,
          },
          startAdornment: <InsertPhotoOutlinedIcon />,
          endAdornment: (
            <InputAdornment position="end">
              <Button
                onClick={handleUpload}
                variant="outlined"
                color={success ? 'success' : 'primary'}
                sx={{ ml: 2 }}
                size="medium"
                disabled={file === null}
              >
                {uploading ? 'Uploading' : success ? 'Success' : 'Upload'}
              </Button>
            </InputAdornment>
          ),
        }}
      />
  );
};

export default FileInput;
```

The file input component allows users to select files, shows upload status, and displays a success message once the upload is complete. We'll implement the core upload logic using a button that triggers the file upload process.

### Uploading Files

We'll complete `handleUpload` function by putting our main upload logic.

```javascript
const handleUpload = async () => {
    if (file === null) return;
    setUploading(true);

    try {
      const { url } = await fetch(
        `${API_BASE_URL}/get-signed-url?filename=${file.name}`
      ).then((res) => res.json());

      await fetch(url, {
        method: 'PUT',
        body: file,
        headers: {
          'Content-Type': file.type,
        },
      });

      const downloadURL = url.split('?')[0];
      // This URL can be used for accessing the uploaded file      
      console.log(downloadURL);
      setSuccess(true);
    } catch (error) {
      console.log(error);
    } finally {
      setUploading(false);
    }
}
```

### File Upload Logic

1. **Get a Signed URL:** The client first sends a request to the server to get a signed URL for uploading the file.
2. **Upload the File:** Using the signed URL, the client sends the file directly to AWS S3.
3. **Get the File URL:** After the file is uploaded, the client can access the file through the URL generated by S3.

That's for Client-side code. With this method, you can easily upload various types of files to S3 from a React client.

Happy coding!
