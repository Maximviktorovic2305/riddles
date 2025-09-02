import { useState, useEffect, useCallback } from 'react';
import { authAPI, AuthResponse } from '@/lib/api';

// Define types
interface User {
  id: number;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
}

interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isLoading: boolean;
  error: string | null;
}

// Initial state
const initialAuthState: AuthState = {
  user: null,
  accessToken: null,
  refreshToken: null,
  isLoading: false,
  error: null,
};

// Helper functions for localStorage
const setAuthData = (data: AuthResponse) => {
  localStorage.setItem('user', JSON.stringify(data.user));
  localStorage.setItem('accessToken', data.access_token);
  localStorage.setItem('refreshToken', data.refresh_token);
};

const getAuthData = (): AuthState => {
  try {
    const user = localStorage.getItem('user');
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');
    
    return {
      user: user ? JSON.parse(user) : null,
      accessToken,
      refreshToken,
      isLoading: false,
      error: null,
    };
  } catch (error) {
    return initialAuthState;
  }
};

const clearAuthData = () => {
  localStorage.removeItem('user');
  localStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
};

// Custom hook
export const useAuth = () => {
  const [authState, setAuthState] = useState<AuthState>(getAuthData());

  // Register function
  const register = useCallback(async (username: string, email: string, password: string) => {
    try {
      setAuthState(prev => ({ ...prev, isLoading: true, error: null }));
      const data = await authAPI.register(username, email, password);
      setAuthData(data);
      setAuthState({
        user: data.user,
        accessToken: data.access_token,
        refreshToken: data.refresh_token,
        isLoading: false,
        error: null,
      });
      return data;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Registration failed';
      setAuthState(prev => ({ ...prev, isLoading: false, error: errorMessage }));
      throw error;
    }
  }, []);

  // Login function
  const login = useCallback(async (email: string, password: string) => {
    try {
      setAuthState(prev => ({ ...prev, isLoading: true, error: null }));
      const data = await authAPI.login(email, password);
      setAuthData(data);
      setAuthState({
        user: data.user,
        accessToken: data.access_token,
        refreshToken: data.refresh_token,
        isLoading: false,
        error: null,
      });
      return data;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Login failed';
      setAuthState(prev => ({ ...prev, isLoading: false, error: errorMessage }));
      throw error;
    }
  }, []);

  // Logout function
  const logout = useCallback(() => {
    clearAuthData();
    setAuthState(initialAuthState);
  }, []);

  // Refresh token function
  const refreshToken = useCallback(async () => {
    if (!authState.refreshToken) {
      logout();
      return;
    }

    try {
      const data = await authAPI.refresh(authState.refreshToken);
      localStorage.setItem('accessToken', data.access_token);
      localStorage.setItem('refreshToken', data.refresh_token);
      setAuthState(prev => ({
        ...prev,
        accessToken: data.access_token,
        refreshToken: data.refresh_token,
        error: null,
      }));
    } catch (error) {
      logout();
      throw error;
    }
  }, [authState.refreshToken, logout]);

  // Check if user is authenticated
  const isAuthenticated = !!authState.accessToken && !!authState.user;

  // Effect to check auth status on mount
  useEffect(() => {
    const data = getAuthData();
    setAuthState(data);
  }, []);

  return {
    ...authState,
    register,
    login,
    logout,
    refreshToken,
    isAuthenticated,
  };
};