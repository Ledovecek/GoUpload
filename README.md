
# GoUpload

Go upload & download archieve files.


## API Reference

#### Upload file

```http
  POST /
```

| form-data key | Type     | Description                |
| :-------- | :------- | :------------------------- |
| file | `file` | **Required**. Archieve file [max 1.1 GB] |

#### Download file

```http
  GET /download/{uuid}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `uuid`      | `string` | **Required**. UUID of uploaded file |

