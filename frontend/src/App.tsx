import React, { useEffect } from 'react';
import { Navigate, Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import AuthLayout from './components/AuthLayout';
import Layout from './components/Layout';
import ProtectedRoute from './components/ProtectedRoute';
import AmortizationPage from './pages/AmortizationPage';
import AssetsPage from './pages/AssetsPage';
import BusinessDebtPage from './pages/BusinessDebtPage';
import CompanyDashboard from './pages/CompanyDashboard';
import DashboardPage from './pages/DashboardPage';
import DepreciationPage from './pages/DepreciationPage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import { useAuthStore } from './store/authStore';

function App() {
  const { user, getMe } = useAuthStore();

  useEffect(() => {
    // Проверяем токен при загрузке приложения
    if (!user) {
      getMe();
    }
  }, [user, getMe]);

  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Routes>
          {/* Публичные маршруты */}
          <Route
            path="/login"
            element={
              user ? (
                <Navigate to="/dashboard" replace />
              ) : (
                <AuthLayout>
                  <LoginPage />
                </AuthLayout>
              )
            }
          />
          <Route
            path="/register"
            element={
              user ? (
                <Navigate to="/dashboard" replace />
              ) : (
                <AuthLayout>
                  <RegisterPage />
                </AuthLayout>
              )
            }
          />

          {/* Защищенные маршруты */}
          <Route
            path="/dashboard"
            element={
              <ProtectedRoute>
                <Layout>
                  <DashboardPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/company/:companyId"
            element={
              <ProtectedRoute>
                <Layout>
                  <CompanyDashboard />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/company/:companyId/assets"
            element={
              <ProtectedRoute>
                <Layout>
                  <AssetsPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/asset/:assetId/amortization"
            element={
              <ProtectedRoute>
                <Layout>
                  <AmortizationPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/asset/:assetId/depreciation"
            element={
              <ProtectedRoute>
                <Layout>
                  <DepreciationPage />
                </Layout>
              </ProtectedRoute>
            }
          />
          <Route
            path="/company/:companyId/business-debt"
            element={
              <ProtectedRoute>
                <Layout>
                  <BusinessDebtPage />
                </Layout>
              </ProtectedRoute>
            }
          />

          {/* Редирект на dashboard по умолчанию */}
          <Route path="/" element={<Navigate to="/dashboard" replace />} />
          
          {/* 404 страница */}
          <Route
            path="*"
            element={
              <div className="min-h-screen flex items-center justify-center">
                <div className="text-center">
                  <h1 className="text-4xl font-bold text-gray-900 mb-4">404</h1>
                  <p className="text-gray-600 mb-8">Страница не найдена</p>
                  <a
                    href="/dashboard"
                    className="bg-primary-600 text-white px-4 py-2 rounded-md hover:bg-primary-700"
                  >
                    Вернуться на главную
                  </a>
                </div>
              </div>
            }
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App; 