import axios from 'axios';
import {
    AmortizationSchedule,
    Asset,
    AssetRequest,
    AuthResponse,
    BusinessDebtSchedule,
    Company,
    CompanyRequest,
    DepreciationSchedule,
    LoginRequest,
    RegisterRequest,
    User,
} from '../types';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// Создаем экземпляр axios
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Интерцептор для добавления токена авторизации
api.interceptors.request.use((config) => {
  const authData = localStorage.getItem('auth-storage');
  if (authData) {
    try {
      const { state } = JSON.parse(authData);
      if (state?.token) {
        config.headers.Authorization = `Bearer ${state.token}`;
      }
    } catch (error) {
      console.error('Error parsing auth data:', error);
    }
  }
  return config;
});

// Интерцептор для обработки ошибок
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Токен недействителен, очищаем localStorage
      localStorage.removeItem('auth-storage');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth API
export const authApi = {
  login: async (credentials: LoginRequest): Promise<AuthResponse> => {
    const response = await api.post('/api/auth/login', credentials);
    return response.data;
  },

  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    const response = await api.post('/api/auth/register', data);
    return response.data;
  },

  getMe: async (): Promise<User> => {
    const response = await api.get('/api/auth/me');
    return response.data;
  },
};

// Companies API
export const companiesApi = {
  getAll: async (): Promise<Company[]> => {
    const response = await api.get('/api/companies');
    return response.data;
  },

  create: async (data: CompanyRequest): Promise<Company> => {
    const response = await api.post('/api/companies', data);
    return response.data;
  },

  update: async (id: string, data: CompanyRequest): Promise<Company> => {
    const response = await api.put(`/api/companies/${id}`, data);
    return response.data;
  },

  delete: async (id: string): Promise<void> => {
    await api.delete(`/api/companies/${id}`);
  },
};

// Assets API
export const assetsApi = {
  getByCompany: async (companyId: string): Promise<Asset[]> => {
    const response = await api.get(`/api/assets?company_id=${companyId}`);
    return response.data;
  },

  create: async (data: AssetRequest): Promise<Asset> => {
    const response = await api.post('/api/assets', data);
    return response.data;
  },

  update: async (id: string, data: AssetRequest): Promise<Asset> => {
    const response = await api.put(`/api/assets/${id}`, data);
    return response.data;
  },

  delete: async (id: string): Promise<void> => {
    await api.delete(`/api/assets/${id}`);
  },
};

// Schedules API
export const schedulesApi = {
  getAmortization: async (assetId: string): Promise<AmortizationSchedule> => {
    const response = await api.get(`/api/schedules/amortization/${assetId}`);
    return response.data;
  },

  getDepreciation: async (assetId: string): Promise<DepreciationSchedule> => {
    const response = await api.get(`/api/schedules/depreciation/${assetId}`);
    return response.data;
  },

  getBusinessDebt: async (companyId: string): Promise<BusinessDebtSchedule> => {
    const response = await api.get(`/api/schedules/business-debt/${companyId}`);
    return response.data;
  },
};

export default api; 