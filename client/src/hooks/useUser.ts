import { useQuery } from '@tanstack/react-query';
import { userAPI, User, UserStatsResponse } from '@/lib/api';
import { useAuth } from './useAuth';

// Query keys
const queryKeys = {
  profile: ['profile'] as const,
  stats: ['stats'] as const,
};

// Custom hooks for user data
export const useProfile = () => {
  const { accessToken, isAuthenticated } = useAuth();
  
  return useQuery<User, Error>({
    queryKey: queryKeys.profile,
    queryFn: () => {
      if (!accessToken) throw new Error('Not authenticated');
      return userAPI.getProfile(accessToken);
    },
    enabled: isAuthenticated && !!accessToken,
  });
};

export const useUserStats = () => {
  const { accessToken, isAuthenticated } = useAuth();
  
  return useQuery<UserStatsResponse, Error>({
    queryKey: queryKeys.stats,
    queryFn: () => {
      if (!accessToken) throw new Error('Not authenticated');
      return userAPI.getUserStats(accessToken);
    },
    enabled: isAuthenticated && !!accessToken,
  });
};