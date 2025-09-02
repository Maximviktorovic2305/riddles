import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { 
  riddlesAPI, 
  dailyRiddleAPI, 
  favoritesAPI, 
  ratingsAPI,
  RiddleWithProgress,
  DailyRiddle,
  CheckAnswerResponse
} from '@/lib/api';
import { useAuth } from './useAuth';

// Query keys
const queryKeys = {
  allRiddles: ['riddles'] as const,
  riddle: (id: number) => ['riddle', id] as const,
  dailyRiddle: ['dailyRiddle'] as const,
};

// Custom hooks for riddles
export const useRiddles = () => {
  return useQuery<RiddleWithProgress[], Error>({
    queryKey: queryKeys.allRiddles,
    queryFn: riddlesAPI.getAllRiddles,
  });
};

export const useRiddle = (id: number) => {
  return useQuery<RiddleWithProgress, Error>({
    queryKey: queryKeys.riddle(id),
    queryFn: () => riddlesAPI.getRiddleById(id),
  });
};

export const useDailyRiddle = () => {
  return useQuery<DailyRiddle, Error>({
    queryKey: queryKeys.dailyRiddle,
    queryFn: dailyRiddleAPI.getTodayRiddle,
  });
};

export const useCheckAnswer = () => {
  const queryClient = useQueryClient();
  
  return useMutation<CheckAnswerResponse, Error, { riddleId: number; answer: string }>({
    mutationFn: ({ riddleId, answer }) => riddlesAPI.checkAnswer(riddleId, answer),
    onSuccess: (_, variables) => {
      // Invalidate and refetch the riddle to update progress
      queryClient.invalidateQueries({ queryKey: queryKeys.riddle(variables.riddleId) });
    },
  });
};

// Custom hooks for favorites
export const useAddFavorite = () => {
  const { accessToken } = useAuth();
  const queryClient = useQueryClient();
  
  return useMutation<void, Error, number>({
    mutationFn: (riddleId) => {
      if (!accessToken) throw new Error('Not authenticated');
      return favoritesAPI.addFavorite(riddleId, accessToken);
    },
    onSuccess: () => {
      // Invalidate and refetch riddles to update favorite status
      queryClient.invalidateQueries({ queryKey: queryKeys.allRiddles });
    },
  });
};

export const useRemoveFavorite = () => {
  const { accessToken } = useAuth();
  const queryClient = useQueryClient();
  
  return useMutation<void, Error, number>({
    mutationFn: (riddleId) => {
      if (!accessToken) throw new Error('Not authenticated');
      return favoritesAPI.removeFavorite(riddleId, accessToken);
    },
    onSuccess: () => {
      // Invalidate and refetch riddles to update favorite status
      queryClient.invalidateQueries({ queryKey: queryKeys.allRiddles });
    },
  });
};

// Custom hooks for ratings
export const useRateRiddle = () => {
  const { accessToken } = useAuth();
  const queryClient = useQueryClient();
  
  return useMutation<void, Error, { riddleId: number; rating: number }>({
    mutationFn: ({ riddleId, rating }) => {
      if (!accessToken) throw new Error('Not authenticated');
      return ratingsAPI.rateRiddle(riddleId, rating, accessToken);
    },
    onSuccess: (_, variables) => {
      // Invalidate and refetch riddles to update ratings
      queryClient.invalidateQueries({ queryKey: queryKeys.allRiddles });
      queryClient.invalidateQueries({ queryKey: queryKeys.riddle(variables.riddleId) });
    },
  });
};

export const useRemoveRating = () => {
  const { accessToken } = useAuth();
  const queryClient = useQueryClient();
  
  return useMutation<void, Error, number>({
    mutationFn: (riddleId) => {
      if (!accessToken) throw new Error('Not authenticated');
      return ratingsAPI.removeRating(riddleId, accessToken);
    },
    onSuccess: (_, riddleId) => {
      // Invalidate and refetch riddles to update ratings
      queryClient.invalidateQueries({ queryKey: queryKeys.allRiddles });
      queryClient.invalidateQueries({ queryKey: queryKeys.riddle(riddleId) });
    },
  });
};