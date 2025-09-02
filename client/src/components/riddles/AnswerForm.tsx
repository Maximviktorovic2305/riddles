'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { useCheckAnswer } from '@/hooks/useRiddles';
import { useAuth } from '@/hooks/useAuth';
import { CheckAnswerResponse } from '@/lib/api';

const formSchema = z.object({
  answer: z.string().min(1, { message: 'Введите ответ' }),
});

interface AnswerFormProps {
  riddleId: number;
  onAnswerChecked?: (result: CheckAnswerResponse) => void;
}

export function AnswerForm({ riddleId, onAnswerChecked }: AnswerFormProps) {
  const [result, setResult] = useState<CheckAnswerResponse | null>(null);
  const { isAuthenticated } = useAuth();
  const { mutate: checkAnswer, isPending } = useCheckAnswer();
  
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      answer: '',
    },
  });

  // If user is not authenticated, show login prompt
  if (!isAuthenticated) {
    return (
      <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-4">
        <p className="text-yellow-800">
          Чтобы отправить ответ, пожалуйста,{" "}
          <a href="/login" className="font-medium text-yellow-700 underline">
            войдите
          </a>{" "}
          или{" "}
          <a href="/register" className="font-medium text-yellow-700 underline">
            зарегистрируйтесь
          </a>
          .
        </p>
      </div>
    );
  }

  function onSubmit(values: z.infer<typeof formSchema>) {
    checkAnswer(
      { riddleId, answer: values.answer },
      {
        onSuccess: (data) => {
          setResult(data);
          if (onAnswerChecked) {
            onAnswerChecked(data);
          }
        },
        onError: (error) => {
          console.error('Error checking answer:', error);
          setResult({
            correct: false,
            message: 'Ошибка при проверке ответа. Попробуйте еще раз.',
          });
        },
      }
    );
  }

  return (
    <div className="space-y-4">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="flex gap-2">
          <FormField
            control={form.control}
            name="answer"
            render={({ field }) => (
              <FormItem className="flex-1">
                <FormControl>
                  <Input 
                    placeholder="Введите ваш ответ..." 
                    {...field}
                    disabled={isPending || !!result?.correct}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button 
            type="submit" 
            disabled={isPending || !!result?.correct}
          >
            {isPending ? 'Проверка...' : 'Ответить'}
          </Button>
        </form>
      </Form>

      {result && (
        <div className={`p-4 rounded-lg ${result.correct ? 'bg-green-50 border border-green-200' : 'bg-red-50 border border-red-200'}`}>
          <p className={`font-medium ${result.correct ? 'text-green-800' : 'text-red-800'}`}>
            {result.message}
          </p>
          {result.correct && (
            <p className="mt-2 text-sm text-green-700">
              Поздравляем! Вы решили эту загадку.
            </p>
          )}
        </div>
      )}
    </div>
  );
}