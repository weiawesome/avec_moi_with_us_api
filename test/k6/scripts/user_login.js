import http from 'k6/http';
import { sleep,check } from 'k6';
const config = require("../config.js").config

export default function () {
    const url = config.baseUrl+config.loginRoute;
    const payload = {
        mail:config.testMail,
        password:config.testPassword,
    };

    const response = http.post(url, JSON.stringify(payload), { headers: config.header });
    check(response, {
        'Status is 200': (res) => res.status === 200,
    });
    check(response, {
        'Response body has token property': (res) => JSON.parse(res.body).hasOwnProperty('token'),
    });


    sleep(1);
}