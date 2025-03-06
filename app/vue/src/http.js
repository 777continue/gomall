import axios from 'axios';
import router from './router';

const http = axios.create();

http.interceptors.request.use(config => {
    const token = localStorage.getItem('jwtToken');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

http.interceptors.response.use(
    response => response,
    error => {
        if (error.response && error.response.status === 401) {
            localStorage.removeItem('jwtToken');
            router.push('/sign-in');
        }
        return Promise.reject(error);
    }
);

export default http;