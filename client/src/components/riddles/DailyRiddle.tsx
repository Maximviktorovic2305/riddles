'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { DailyRiddle as DailyRiddleType } from '@/lib/api';
import { AnswerForm } from '@/components/riddles/AnswerForm';

interface DailyRiddleProps {
  riddle: DailyRiddleType;
}

export function DailyRiddle({ riddle }: DailyRiddleProps) {
  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'easy': return 'bg-green-100 text-green-800';
      case 'medium': return 'bg-yellow-100 text-yellow-800';
      case 'hard': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <Card className="w-full">
      <CardHeader>
        <div className="flex justify-between items-start">
          <CardTitle className="text-xl">Загадка дня</CardTitle>
          <Badge variant="default">Сегодня</Badge>
        </div>
        <CardTitle className="text-lg">{riddle.riddle.title}</CardTitle>
        <div className="flex space-x-2">
          <Badge variant="secondary">{riddle.riddle.category.name}</Badge>
          <Badge className={getDifficultyColor(riddle.riddle.difficulty)}>
            {riddle.riddle.difficulty === 'easy' && 'Легко'}
            {riddle.riddle.difficulty === 'medium' && 'Средне'}
            {riddle.riddle.difficulty === 'hard' && 'Сложно'}
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent>
        <p className="text-gray-700 mb-4">{riddle.riddle.description}</p>
        <AnswerForm riddleId={riddle.riddle.id} />
      </CardContent>
    </Card>
  );
}