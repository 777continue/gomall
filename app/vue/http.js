import axios from 'axios';
import router from './router';

const http = axios.create({
    baseURL: process.env.VUE_APP_API_BASE_URL || '/api'
});

http.interceptors.request.use(config => {
    const token = localStorage.getItem('JWTToken');
    if (token) {
        config.headers.Authorization = `jwt ${token}`;
    }
    return config;
});

http.interceptors.response.use(
    response => response,
    error => {
        if (error.response.status === 401) {
            localStorage.removeItem('jwtToken');
            router.push('/sign-in');
        }
        return Promise.reject(error);
    }
);

export default http;