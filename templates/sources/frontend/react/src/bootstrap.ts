declare global {
    interface Window {
        axios: any;
    }
}

import axios from 'axios';
window.axios = axios;

const cookies = document.cookie.split(';');
let csrfToken = '';

for (let i = 0; i < cookies.length; i++) {
    const cookieParts = cookies[i].trim().split('=');
    if (cookieParts[0] === '_csrf') {
        csrfToken = cookieParts[1];
        break;
    }
}

if (csrfToken) {
    axios.defaults.headers.common['X-CSRF-TOKEN'] = csrfToken;
    console.log('CSRF token found: ' + csrfToken);
} else {
    console.error('CSRF token not found: https://inertiajs.com/csrf');
}