import axios from 'axios';
import router from './router';

const http = axios.create();

http.interceptors.request.use(
    config => {
        const token = localStorage.getItem('jwtToken');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    }
);

http.interceptors.response.use(
    response => {
        // 检查是否有新的token
        const newToken = response.headers['x-renewed-token']
        if (newToken) {
            // 更新本地存储的token
            localStorage.setItem('jwtToken', newToken)
        }
        return response
    },
    error => {
        console.log('Error object:', error);
        if (error.response && error.response.status === 401) {
            localStorage.removeItem('jwtToken');
            localStorage.removeItem('admin');
            router.push('/sign-in');
        }
        return Promise.reject(error);
    }
);

export default http;