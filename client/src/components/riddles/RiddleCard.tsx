'use client';

import { useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { 
  ThumbsUp, 
  ThumbsDown, 
  Heart, 
  HeartOff,
  CheckCircle,
  Circle
} from 'lucide-react';
import { RiddleWithProgress } from '@/lib/api';
import { useAuth } from '@/hooks/useAuth';
import { 
  useAddFavorite, 
  useRemoveFavorite, 
  useRateRiddle 
} from '@/hooks/useRiddles';
import { AnswerForm } from '@/components/riddles/AnswerForm';

interface RiddleCardProps {
  riddle: RiddleWithProgress;
  onExpand?: (riddle: RiddleWithProgress) => void;
}

export function RiddleCard({ riddle, onExpand }: RiddleCardProps) {
  const { isAuthenticated } = useAuth();
  const [isExpanded, setIsExpanded] = useState(false);
  
  const { mutate: addFavorite } = useAddFavorite();
  const { mutate: removeFavorite } = useRemoveFavorite();
  const { mutate: rateRiddle } = useRateRiddle();

  const handleToggleFavorite = () => {
    if (!isAuthenticated) {
      // TODO: Show login modal
      return;
    }
    
    if (riddle.is_favorite) {
      removeFavorite(riddle.riddle.id);
    } else {
      addFavorite(riddle.riddle.id);
    }
  };

  const handleRate = (rating: number) => {
    if (!isAuthenticated) {
      // TODO: Show login modal
      return;
    }
    
    rateRiddle({ riddleId: riddle.riddle.id, rating });
  };

  const toggleExpand = () => {
    if (onExpand) {
      onExpand(riddle);
    } else {
      setIsExpanded(!isExpanded);
    }
  };

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'easy': return 'bg-green-100 text-green-800';
      case 'medium': return 'bg-yellow-100 text-yellow-800';
      case 'hard': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <Card className="w-full hover:shadow-md transition-shadow">
      <CardHeader className="pb-2">
        <div className="flex justify-between items-start">
          <CardTitle 
            className="text-lg cursor-pointer hover:underline"
            onClick={toggleExpand}
          >
            {riddle.riddle.title}
          </CardTitle>
          <div className="flex space-x-1">
            {riddle.is_solved ? (
              <CheckCircle className="h-5 w-5 text-green-500" />
            ) : (
              <Circle className="h-5 w-5 text-gray-400" />
            )}
            <Button
              variant="ghost"
              size="icon"
              onClick={handleToggleFavorite}
              className="h-5 w-5"
            >
              {riddle.is_favorite ? (
                <Heart className="h-4 w-4 fill-red-500 text-red-500" />
              ) : (
                <HeartOff className="h-4 w-4" />
              )}
            </Button>
          </div>
        </div>
        <div className="flex justify-between items-center">
          <Badge variant="secondary">{riddle.riddle.category.name}</Badge>
          <Badge className={getDifficultyColor(riddle.riddle.difficulty)}>
            {riddle.riddle.difficulty === 'easy' && 'Легко'}
            {riddle.riddle.difficulty === 'medium' && 'Средне'}
            {riddle.riddle.difficulty === 'hard' && 'Сложно'}
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent>
        <p className="text-sm text-gray-600 mb-4 line-clamp-2">
          {riddle.riddle.description}
        </p>
        
        <div className="flex justify-between items-center">
          <div className="flex space-x-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => handleRate(1)}
              className="h-8 px-2"
            >
              <ThumbsUp className="h-4 w-4 mr-1" />
              <span>{riddle.likes}</span>
            </Button>
            <Button
              variant="ghost"
              size="sm"
              onClick={() => handleRate(-1)}
              className="h-8 px-2"
            >
              <ThumbsDown className="h-4 w-4 mr-1" />
              <span>{riddle.dislikes}</span>
            </Button>
          </div>
          
          <Button 
            variant="outline" 
            size="sm" 
            onClick={toggleExpand}
          >
            {isExpanded ? 'Скрыть' : 'Решить'}
          </Button>
        </div>
        
        {isExpanded && (
          <div className="mt-4 pt-4 border-t space-y-4">
            <p className="text-sm">{riddle.riddle.description}</p>
            <AnswerForm riddleId={riddle.riddle.id} />
          </div>
        )}
      </CardContent>
    </Card>
  );
}