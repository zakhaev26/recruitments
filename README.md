# Starting the Server

## Starting PostgreSQL Container

To start the PostgreSQL container, run the following command:

```bash
docker-compose -f docker/postgres-container.yml up
```
This command will start the PostgreSQL container required for the server to run.

## Starting Go Server Locally

To start the Go server locally, follow these steps:
- Make sure PostgreSQL container is up and running as instructed above
- Navigate to the root directory of the project.
- Run the following command in a new terminal to start the server:

```bash
go run cmd/main/main.go
```

This command will start the Go server locally, and it will be accessible at http://localhost:3000.

## API Routes 

## Register as Applicant

- Request 

```bash
    curl --location '127.0.0.1:3000/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user": {
        "name": "Soubhik Gon",
        "email": "soubhik@mail.com",
        "password": "password",
        "address": "123 Main St",
        "userType": "applicant"
    },
    "profile": {
        "resumeFileAddress": "https://hosted-resume-link",
        "skills": "JavaScript, Python, SQL",
        "education": "Bachelor'\''s in Computer Science",
        "experience": "Software Engineer at XYZ Company",
        "name": "Soubhik Gon",
        "email": "soubhik@mail.com",
        "phone": "+1234567890"
    }
}
'
```

## Register as Admin

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user": {
        "name": "iamaadmin",
        "email": "iamaadmin@mail.com",
        "password": "password",
        "address": "123 Main St",
        "userType": "admin"
    },
    "profile": {
        "resumeFileAddress": "https://hosted-resume-link",
        "skills": "JavaScript, Python, SQL",
        "education": "Bachelor'\''s in Computer Science",
        "experience": "Software Engineer at XYZ Company",
        "name": "Soubhik Gon",
        "email": "soubhik@mail.com",
        "phone": "+1234567890"
    }
}
'
```

## Login as User

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/login' \
--data-raw '{
    "email": "soubhik@mail.com",
    "password": "password"
}'
```

> Response : 

```bash
{
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE5MDE4NDkyLTNkOTQtNDczYi1hM2Q3LTA5YWY2OTgyNTIyZCIsInVzZXJUeXBlIjoiYXBwbGljYW50IiwiUHJvZmlsZUlEIjoiZjVkYTA0Y2EtZmY1Ny00NTBiLTliNTgtYTkyMDEyOWEyZDIyIiwiZXhwIjoxNzEzNzE2NjMwLCJpYXQiOjE3MTM2MzAyMzB9.nUYXtHEpGAohvQq8J-sbgsNLGM79g_0LUjkbZJFOhkc"
}
```

> NOTE : Use this accessToken value in the request's authorization as Bearer <accessToken>

## Upload resume

- Request 
```bash
curl --location '127.0.0.1:3000/api/v1/uploadResume' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE5MDE4NDkyLTNkOTQtNDczYi1hM2Q3LTA5YWY2OTgyNTIyZCIsInVzZXJUeXBlIjoiYXBwbGljYW50IiwiUHJvZmlsZUlEIjoiZjVkYTA0Y2EtZmY1Ny00NTBiLTliNTgtYTkyMDEyOWEyZDIyIiwiZXhwIjoxNzEzNzE2NjMwLCJpYXQiOjE3MTM2MzAyMzB9.nUYXtHEpGAohvQq8J-sbgsNLGM79g_0LUjkbZJFOhkc' \
--header 'Content-Type: text/plain' \
--data '@/home/soubhik/Downloads/soubhik_resume_march.pdf'
```

> Response :

```bash
File uploaded successfully. File saved in the database.
```

> NOTE: This endpoint calls internally the resume parse api (a third-party api) and it takes time to parse the CV.Future wokr would be to make a separate server and do this task asynchronously using a message broker (RabbitMQ/Kafka) to resuce response times to client
(which is currently approximately ~ 20s)

## Login as Admin

- Request
```bash
curl --location '127.0.0.1:3000/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "iamaadmin@mail.com",
    "password": "password"
}'
```

> Response: 
```bash
{
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlmODI0YWE4LTYwYmItNDI2ZC1hN2UyLWFhNTg1MzYzNWQ4MSIsInVzZXJUeXBlIjoiYWRtaW4iLCJQcm9maWxlSUQiOiJkZjUxOTE0YS1iOTEyLTQyZjQtODA2My01NTJhYTFjYzEyOTgiLCJleHAiOjE3MTM3MTY5NjEsImlhdCI6MTcxMzYzMDU2MX0.3737cbdLOm95W7w-NHtA8en8QyHZNRdI90Cf5Cj36B4"
}
```

## Create a Job Posting

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/admin/job' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlmODI0YWE4LTYwYmItNDI2ZC1hN2UyLWFhNTg1MzYzNWQ4MSIsInVzZXJUeXBlIjoiYWRtaW4iLCJQcm9maWxlSUQiOiJkZjUxOTE0YS1iOTEyLTQyZjQtODA2My01NTJhYTFjYzEyOTgiLCJleHAiOjE3MTM3MTY5NjEsImlhdCI6MTcxMzYzMDU2MX0.3737cbdLOm95W7w-NHtA8en8QyHZNRdI90Cf5Cj36B4' \
--header 'Content-Type: application/json' \
--data-raw '{
  "title": "Go Engineer",
  "description": "We are looking for a Go software engineer to join our team.",
  "companyName": "Synergy Labs",
  "email": "iamaadmin@mail.com"
}
'
```

> Response:
```bash
{
    "CreatedAt": "2024-04-20T22:01:09.006565678+05:30",
    "UpdatedAt": "2024-04-20T22:01:09.006565678+05:30",
    "DeletedAt": null,
    "ID": "1cd2d0e3-a565-4479-8d24-43152fae428b",
    "Title": "Go Engineer",
    "Description": "We are looking for a Go software engineer to join our team.",
    "PostedOn": "2024-04-20T22:01:09.003244197+05:30",
    "TotalApplications": 0,
    "CompanyName": "Synergy Labs",
    "PostedByID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
    "PostedBy": {
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "ID": "00000000-0000-0000-0000-000000000000",
        "Name": "",
        "Email": "",
        "Address": "",
        "UserType": "",
        "Password": "",
        "ProfileHeadline": "",
        "ProfileID": "00000000-0000-0000-0000-000000000000",
        "Profile": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "00000000-0000-0000-0000-000000000000",
            "ResumeFileAddress": "",
            "Skills": "",
            "Education": "",
            "Experience": "",
            "Name": "",
            "Email": "",
            "Phone": ""
        }
    }
}
```

## Get all job listings

- Request

```bash
curl --location '127.0.0.1:3000/api/v1/jobs' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjlmODI0YWE4LTYwYmItNDI2ZC1hN2UyLWFhNTg1MzYzNWQ4MSIsInVzZXJUeXBlIjoiYWRtaW4iLCJQcm9maWxlSUQiOiJkZjUxOTE0YS1iOTEyLTQyZjQtODA2My01NTJhYTFjYzEyOTgiLCJleHAiOjE3MTM3MTY5NjEsImlhdCI6MTcxMzYzMDU2MX0.3737cbdLOm95W7w-NHtA8en8QyHZNRdI90Cf5Cj36B4'
```

- Respone
```bash
[
    {
        "CreatedAt": "2024-04-21T00:31:09.006565+08:00",
        "UpdatedAt": "2024-04-21T00:31:09.006565+08:00",
        "DeletedAt": null,
        "ID": "1cd2d0e3-a565-4479-8d24-43152fae428b",
        "Title": "Go Engineer",
        "Description": "We are looking for a Go software engineer to join our team.",
        "PostedOn": "2024-04-21T00:31:09.003244+08:00",
        "TotalApplications": 0,
        "CompanyName": "Synergy Labs",
        "PostedByID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
        "PostedBy": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
            "Name": "iamaadmin",
            "Email": "iamaadmin@mail.com",
            "Address": "",
            "UserType": "admin",
            "Password": "",
            "ProfileHeadline": "",
            "ProfileID": "df51914a-b912-42f4-8063-552aa1cc1298",
            "Profile": {
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "ID": "df51914a-b912-42f4-8063-552aa1cc1298",
                "ResumeFileAddress": "",
                "Skills": "JavaScript, Python, SQL",
                "Education": "Bachelor's in Computer Science",
                "Experience": "",
                "Name": "",
                "Email": "",
                "Phone": ""
            }
        }
    }
]
```

> NOTE: adding more jobs will result in populating of the above array,as an example here's only one.

## Apply for a job

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/jobs/apply?job_id=1cd2d0e3-a565-4479-8d24-43152fae428b' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE5MDE4NDkyLTNkOTQtNDczYi1hM2Q3LTA5YWY2OTgyNTIyZCIsInVzZXJUeXBlIjoiYXBwbGljYW50IiwiUHJvZmlsZUlEIjoiZjVkYTA0Y2EtZmY1Ny00NTBiLTliNTgtYTkyMDEyOWEyZDIyIiwiZXhwIjoxNzEzNzE3MzI5LCJpYXQiOjE3MTM2MzA5Mjl9.ewlESItndgogvFS1LAvSMIcey2pBJvMI0Do1CaaAfk4'
```

- Response 
```bash
Job application submitted successfully
```

## Get all System Users

- Request

```bash 
[
    {
        "CreatedAt": "2024-04-21T00:20:32.384875+08:00",
        "UpdatedAt": "2024-04-21T00:20:32.384875+08:00",
        "DeletedAt": null,
        "ID": "19018492-3d94-473b-a3d7-09af6982522d",
        "Name": "Soubhik Gon",
        "Email": "soubhik@mail.com",
        "Address": "123 Main St",
        "UserType": "applicant",
        "Password": "",
        "ProfileHeadline": "",
        "ProfileID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
        "Profile": {
            "CreatedAt": "2024-04-21T00:20:32.358477+08:00",
            "UpdatedAt": "2024-04-21T00:25:37.331559+08:00",
            "DeletedAt": null,
            "ID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
            "ResumeFileAddress": "https://hosted-resume-link",
            "Skills": "C++, Pandas, Database, Watchdog, Rest, International, System, Erp, Audio, Flask, Github, Cloud, Python, R, Javascript, Linux, Video, Information technology, Ruby, Matplotlib, Zookeeper, Os, Aws, Content, C, Docker, Email, Apis, Html, Css",
            "Education": "International Institute Of Information Technology, Bhubaneswar B.Tech Information Technology | GPA (January 2010, January 2022, November 2022, January 2024, July 2026); International Institute Of Information Technology (November 2022); Delhi Public School (January 2010); Delhi Public School Kalinga (January 2010)",
            "Experience": "Devops at AWS",
            "Name": "Soubhik Kumar Gon",
            "Email": "soubhik@mail.com",
            "Phone": "2010 - 2022"
        }
    },
    {
        "CreatedAt": "2024-04-21T00:20:32.413769+08:00",
        "UpdatedAt": "2024-04-21T00:20:32.413769+08:00",
        "DeletedAt": null,
        "ID": "313de22a-5249-487d-b797-7ac0a2c46245",
        "Name": "Soubhik Gon",
        "Email": "soubhik@mail.com",
        "Address": "123 Main St",
        "UserType": "applicant",
        "Password": "",
        "ProfileHeadline": "",
        "ProfileID": "919f5d72-41a0-4154-a762-7ab156ed1c0d",
        "Profile": {
            "CreatedAt": "2024-04-21T00:20:32.39679+08:00",
            "UpdatedAt": "2024-04-21T00:20:32.39679+08:00",
            "DeletedAt": null,
            "ID": "919f5d72-41a0-4154-a762-7ab156ed1c0d",
            "ResumeFileAddress": "https://hosted-resume-link",
            "Skills": "JavaScript, Python, SQL",
            "Education": "Bachelor's in Computer Science",
            "Experience": "Software Engineer at XYZ Company",
            "Name": "Soubhik Gon",
            "Email": "soubhik@mail.com",
            "Phone": "+1234567890"
        }
    },
    {
        "CreatedAt": "2024-04-21T00:23:04.873688+08:00",
        "UpdatedAt": "2024-04-21T00:23:04.873688+08:00",
        "DeletedAt": null,
        "ID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
        "Name": "iamaadmin",
        "Email": "iamaadmin@mail.com",
        "Address": "123 Main St",
        "UserType": "admin",
        "Password": "",
        "ProfileHeadline": "",
        "ProfileID": "df51914a-b912-42f4-8063-552aa1cc1298",
        "Profile": {
            "CreatedAt": "2024-04-21T00:23:04.851544+08:00",
            "UpdatedAt": "2024-04-21T00:23:04.851544+08:00",
            "DeletedAt": null,
            "ID": "df51914a-b912-42f4-8063-552aa1cc1298",
            "ResumeFileAddress": "https://hosted-resume-link",
            "Skills": "JavaScript, Python, SQL",
            "Education": "Bachelor's in Computer Science",
            "Experience": "Software Engineer at XYZ Company",
            "Name": "Soubhik Gon",
            "Email": "soubhik@mail.com",
            "Phone": "+1234567890"
        }
    }
]
```

## Get a job detail:

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/admin/job?job_id=1cd2d0e3-a565-4479-8d24-43152fae428b' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImI3MzkxOGQxLThiYzQtNDY1Yi1hZGUwLWU3OTU0MWJkZmM2OSIsInVzZXJUeXBlIjoiYWRtaW4iLCJQcm9maWxlSUQiOiJkZGMzMmFiNi02YmY5LTQ5YzctOGZkYi04N2Q3Yzg1ZjYwNjQiLCJleHAiOjE3MTM2ODM3NjcsImlhdCI6MTcxMzU5NzM2N30.pe5X8xV9AuXvZ0TCwF0FMgLZyVj4d992KAWGEdIzN44' \
--data ''
```

- Response
```bash
{
    "jobDetails": {
        "CreatedAt": "2024-04-21T00:31:09.006565+08:00",
        "UpdatedAt": "2024-04-21T00:31:09.006565+08:00",
        "DeletedAt": null,
        "ID": "1cd2d0e3-a565-4479-8d24-43152fae428b",
        "Title": "Go Engineer",
        "Description": "We are looking for a Go software engineer to join our team.",
        "PostedOn": "2024-04-21T00:31:09.003244+08:00",
        "TotalApplications": 0,
        "CompanyName": "Synergy Labs",
        "PostedByID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
        "PostedBy": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
            "Name": "iamaadmin",
            "Email": "iamaadmin@mail.com",
            "Address": "123 Main St",
            "UserType": "admin",
            "Password": "",
            "ProfileHeadline": "",
            "ProfileID": "df51914a-b912-42f4-8063-552aa1cc1298",
            "Profile": {
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "ID": "00000000-0000-0000-0000-000000000000",
                "ResumeFileAddress": "",
                "Skills": "",
                "Education": "",
                "Experience": "",
                "Name": "",
                "Email": "",
                "Phone": ""
            }
        }
    },
    "applicants": [
        {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "19018492-3d94-473b-a3d7-09af6982522d",
            "Name": "Soubhik Gon",
            "Email": "soubhik@mail.com",
            "Address": "123 Main St",
            "UserType": "applicant",
            "Password": "",
            "ProfileHeadline": "",
            "ProfileID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
            "Profile": {
                "CreatedAt": "2024-04-21T00:20:32.358477+08:00",
                "UpdatedAt": "2024-04-21T00:25:37.331559+08:00",
                "DeletedAt": null,
                "ID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
                "ResumeFileAddress": "https://hosted-resume-link",
                "Skills": "C++, Pandas, Database, Watchdog, Rest, International, System, Erp, Audio, Flask, Github, Cloud, Python, R, Javascript, Linux, Video, Information technology, Ruby, Matplotlib, Zookeeper, Os, Aws, Content, C, Docker, Email, Apis, Html, Css",
                "Education": "International Institute Of Information Technology, Bhubaneswar B.Tech Information Technology | GPA (January 2010, January 2022, November 2022, January 2024, July 2026); International Institute Of Information Technology (November 2022); Delhi Public School (January 2010); Delhi Public School Kalinga (January 2010)",
                "Experience": "Devops at AWS",
                "Name": "Soubhik Kumar Gon",
                "Email": "soubhik@mail.com",
                "Phone": "2010 - 2022"
            }
        }
    ]
}
```

## Get applicant information

- Request 

```bash
curl --location '127.0.0.1:3000/api/v1/admin/applicant?applicant_id=19018492-3d94-473b-a3d7-09af6982522d' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImI3MzkxOGQxLThiYzQtNDY1Yi1hZGUwLWU3OTU0MWJkZmM2OSIsInVzZXJUeXBlIjoiYWRtaW4iLCJQcm9maWxlSUQiOiJkZGMzMmFiNi02YmY5LTQ5YzctOGZkYi04N2Q3Yzg1ZjYwNjQiLCJleHAiOjE3MTM2ODM3NjcsImlhdCI6MTcxMzU5NzM2N30.pe5X8xV9AuXvZ0TCwF0FMgLZyVj4d992KAWGEdIzN44'
```

- Response

```bash
{
    "jobDetails": {
        "CreatedAt": "2024-04-21T00:31:09.006565+08:00",
        "UpdatedAt": "2024-04-21T00:31:09.006565+08:00",
        "DeletedAt": null,
        "ID": "1cd2d0e3-a565-4479-8d24-43152fae428b",
        "Title": "Go Engineer",
        "Description": "We are looking for a Go software engineer to join our team.",
        "PostedOn": "2024-04-21T00:31:09.003244+08:00",
        "TotalApplications": 0,
        "CompanyName": "Synergy Labs",
        "PostedByID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
        "PostedBy": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "9f824aa8-60bb-426d-a7e2-aa5853635d81",
            "Name": "iamaadmin",
            "Email": "iamaadmin@mail.com",
            "Address": "123 Main St",
            "UserType": "admin",
            "Password": "",
            "ProfileHeadline": "",
            "ProfileID": "df51914a-b912-42f4-8063-552aa1cc1298",
            "Profile": {
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "ID": "00000000-0000-0000-0000-000000000000",
                "ResumeFileAddress": "",
                "Skills": "",
                "Education": "",
                "Experience": "",
                "Name": "",
                "Email": "",
                "Phone": ""
            }
        }
    },
    "applicants": [
        {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": "19018492-3d94-473b-a3d7-09af6982522d",
            "Name": "Soubhik Gon",
            "Email": "soubhik@mail.com",
            "Address": "123 Main St",
            "UserType": "applicant",
            "Password": "",
            "ProfileHeadline": "",
            "ProfileID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
            "Profile": {
                "CreatedAt": "2024-04-21T00:20:32.358477+08:00",
                "UpdatedAt": "2024-04-21T00:25:37.331559+08:00",
                "DeletedAt": null,
                "ID": "f5da04ca-ff57-450b-9b58-a920129a2d22",
                "ResumeFileAddress": "https://hosted-resume-link",
                "Skills": "C++, Pandas, Database, Watchdog, Rest, International, System, Erp, Audio, Flask, Github, Cloud, Python, R, Javascript, Linux, Video, Information technology, Ruby, Matplotlib, Zookeeper, Os, Aws, Content, C, Docker, Email, Apis, Html, Css",
                "Education": "International Institute Of Information Technology, Bhubaneswar B.Tech Information Technology | GPA (January 2010, January 2022, November 2022, January 2024, July 2026); International Institute Of Information Technology (November 2022); Delhi Public School (January 2010); Delhi Public School Kalinga (January 2010)",
                "Experience": "Devops at AWS",
                "Name": "Soubhik Kumar Gon",
                "Email": "soubhik@mail.com",
                "Phone": "2010 - 2022"
            }
        }
    ]
}
```

