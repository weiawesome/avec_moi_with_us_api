import http from 'k6/http';
import { sleep,check } from 'k6';
const config = require("../config.js").config
function generateRandomString(length) {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    let result = "";
    for (let i = 0; i < length; i++) {
        const randomIndex = Math.floor(Math.random() * charset.length);
        result += charset.charAt(randomIndex);
    }
    return result;
}

export default function () {
    const randomPage = Math.floor(Math.random() * 5) + 1;
    const randomLength = Math.floor(Math.random() * 5) + 1;
    const url = config.baseUrl+config.movieRoute+config.pageParameter+randomPage+config.contentParameter+generateRandomString(randomLength);

    const response = http.get(url);

    check(response, {
        'Status is 200': (res) => res.status === 200,
    });
    check(response, {
        'Response body has current_page property': (res) => JSON.parse(res.body).hasOwnProperty('current_page'),
    });
    check(response, {
        'Response body has total_pages property': (res) => JSON.parse(res.body).hasOwnProperty('total_pages'),
    });
    check(response, {
        'Response body has movies property': (res) => JSON.parse(res.body).hasOwnProperty('movies'),
    });
    sleep(1);
}