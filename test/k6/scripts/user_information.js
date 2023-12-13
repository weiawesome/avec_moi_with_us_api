import http from 'k6/http';
import { sleep,check } from 'k6';
const config = require("../config.js").config


export default function () {
    const url = config.baseUrl+config.informationRoute;

    const response = http.get(url, { headers: config.jwtHeader });

    check(response, {
        'Status is 200': (res) => res.status === 200,
    });
    check(response, {
        'Response body has mail property': (res) => JSON.parse(res.body).hasOwnProperty('mail'),
    });
    check(response, {
        'Response body has name property': (res) => JSON.parse(res.body).hasOwnProperty('name'),
    });
    check(response, {
        'Response body has gender property': (res) => JSON.parse(res.body).hasOwnProperty('gender'),
    });
    sleep(1);
}