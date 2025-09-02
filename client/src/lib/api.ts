// API client for communicating with the backend
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api';

// Types
export interface User {
  id: number;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface Category {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

export interface Riddle {
  id: number;
  title: string;
  description: string;
  answer: string;
  category_id: number;
  category: Category;
  difficulty: string; // 'easy' | 'medium' | 'hard'
  created_at: string;
  updated_at: string;
}

export interface RiddleWithProgress {
  riddle: Riddle;
  is_solved: boolean;
  is_favorite: boolean;
  user_rating: number; // -1, 0, or 1
  likes: number;
  dislikes: number;
}

export interface DailyRiddle {
  id: number;
  riddle_id: number;
  riddle: Riddle;
  featured_date: string;
  created_at: string;
}

export interface AuthResponse {
  user: User;
  access_token: string;
  refresh_token: string;
}

export interface RefreshResponse {
  access_token: string;
  refresh_token: string;
}

export interface CheckAnswerResponse {
  correct: boolean;
  message: string;
}

export interface UserStatsResponse {
  total_riddles: number;
  solved_riddles: number;
  success_rate: number;
}

// Auth API
export const authAPI = {
  register: async (username: string, email: string, password: string): Promise<AuthResponse> => {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, email, password }),
    });

    if (!response.ok) {
      throw new Error('Failed to register');
    }

    return response.json();
  },

  login: async (email: string, password: string): Promise<AuthResponse> => {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      throw new Error('Failed to login');
    }

    return response.json();
  },

  refresh: async (refreshToken: string): Promise<RefreshResponse> => {
    const response = await fetch(`${API_BASE_URL}/auth/refresh`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ refresh_token: refreshToken }),
    });

    if (!response.ok) {
      throw new Error('Failed to refresh token');
    }

    return response.json();
  },
};

// Riddles API
export const riddlesAPI = {
  getAllRiddles: async (): Promise<RiddleWithProgress[]> => {
    const response = await fetch(`${API_BASE_URL}/riddles`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch riddles');
    }

    return response.json();
  },

  getRiddleById: async (id: number): Promise<RiddleWithProgress> => {
    const response = await fetch(`${API_BASE_URL}/riddles/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch riddle');
    }

    return response.json();
  },

  checkAnswer: async (id: number, answer: string): Promise<CheckAnswerResponse> => {
    const response = await fetch(`${API_BASE_URL}/riddles/${id}/answer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ answer }),
    });

    if (!response.ok) {
      throw new Error('Failed to check answer');
    }

    return response.json();
  },
};

// Daily Riddle API
export const dailyRiddleAPI = {
  getTodayRiddle: async (): Promise<DailyRiddle> => {
    const response = await fetch(`${API_BASE_URL}/daily-riddle`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch today\'s riddle');
    }

    return response.json();
  },
};

// User API (requires authentication)
export const userAPI = {
  getProfile: async (accessToken: string): Promise<User> => {
    const response = await fetch(`${API_BASE_URL}/users/profile`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch profile');
    }

    return response.json();
  },

  getUserStats: async (accessToken: string): Promise<UserStatsResponse> => {
    const response = await fetch(`${API_BASE_URL}/users/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch user stats');
    }

    return response.json();
  },
};

// Favorites API (requires authentication)
export const favoritesAPI = {
  addFavorite: async (riddleId: number, accessToken: string): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/favorites/${riddleId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to add favorite');
    }
  },

  removeFavorite: async (riddleId: number, accessToken: string): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/favorites/${riddleId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to remove favorite');
    }
  },
};

// Ratings API (requires authentication)
export const ratingsAPI = {
  rateRiddle: async (riddleId: number, rating: number, accessToken: string): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/ratings/${riddleId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
      body: JSON.stringify({ rating }),
    });

    if (!response.ok) {
      throw new Error('Failed to rate riddle');
    }
  },

  removeRating: async (riddleId: number, accessToken: string): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/ratings/${riddleId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to remove rating');
    }
  },
};