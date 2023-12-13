const jwtToken="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYWNjZXNzIiwiZnJlc2giOmZhbHNlLCJjc3JmIjoiMzQ1NTgwNWUtOGRlMi00YWQzLWIyMjUtZmI1YWJkNWRlMzgyIiwiaXNzIjoiRGVmYXVsdElzc3VlciIsInN1YiI6IlRjd2VlZWlAZ21haWwuY29tIiwiZXhwIjoxNzAzNjg1NTMyLCJuYmYiOjE3MDI0NzU5MzIsImlhdCI6MTcwMjQ3NTkzMiwianRpIjoiNzk1YjEwYjQtNGE1Mi00MDM4LWFhNDEtZTg1MGMzNTE4NWFlIn0.Ra-DmMPmoPQHhCJfM1d6qd1Hmio2fzXAwaqPb_e9WT8";
export const config = {
    baseUrl: "https://tcweeei.study-savvy.com",
    loginRoute: "/api/v1/user/login",
    informationRoute:"/api/v1/user",
    movieRoute: "/api/v1/movie",
    pageParameter: "?page=",
    movieSearchRoute: "/api/v1/movie/search",
    contentParameter:"&content=",
    header:{ 'Content-Type': 'application/json' },
    jwtHeader:{'Content-Type': 'application/json', 'Authorization': `Bearer ${jwtToken}`},
    testMail: "LoadTest@gmail.com",
    testPassword:"LoadTestPassword",
    username: "your_username",
    password: "your_password",
};
