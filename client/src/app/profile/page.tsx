"use client";

import { useProfile, useUserStats } from "@/hooks/useUser";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

export default function ProfilePage() {
  const { data: user, isLoading: isUserLoading, error: userError } = useProfile();
  const { data: stats, isLoading: isStatsLoading, error: statsError } = useUserStats();

  // Show loading state
  if (isUserLoading || isStatsLoading) {
    return <div className="flex justify-center items-center h-64">Загрузка...</div>;
  }

  // Show error state
  if (userError || statsError) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-red-500">Ошибка загрузки данных</div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">Профиль</h1>

      {/* User Info */}
      <Card>
        <CardHeader>
          <CardTitle>Информация о пользователе</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2">
          {user && (
            <>
              <div className="flex items-center space-x-4">
                <div className="h-16 w-16 rounded-full bg-gray-200 flex items-center justify-center">
                  <span className="text-2xl font-bold text-gray-600">
                    {user.username.charAt(0).toUpperCase()}
                  </span>
                </div>
                <div>
                  <h2 className="text-xl font-semibold">{user.username}</h2>
                  <p className="text-gray-600">{user.email}</p>
                </div>
              </div>
              <div className="pt-4">
                <p className="text-sm text-gray-500">
                  Участник с {new Date(user.created_at).toLocaleDateString("ru-RU")}
                </p>
              </div>
            </>
          )}
        </CardContent>
      </Card>

      {/* Stats */}
      <Card>
        <CardHeader>
          <CardTitle>Статистика</CardTitle>
        </CardHeader>
        <CardContent>
          {stats ? (
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div className="bg-blue-50 p-4 rounded-lg text-center">
                <div className="text-2xl font-bold">{stats.total_riddles}</div>
                <div className="text-sm text-gray-600">Всего загадок</div>
              </div>
              <div className="bg-green-50 p-4 rounded-lg text-center">
                <div className="text-2xl font-bold">{stats.solved_riddles}</div>
                <div className="text-sm text-gray-600">Решено</div>
              </div>
              <div className="bg-purple-50 p-4 rounded-lg text-center">
                <div className="text-2xl font-bold">{stats.success_rate}%</div>
                <div className="text-sm text-gray-600">Успешность</div>
              </div>
            </div>
          ) : (
            <div className="text-center py-4 text-gray-500">
              Статистика недоступна
            </div>
          )}
        </CardContent>
      </Card>

      {/* Achievements */}
      <Card>
        <CardHeader>
          <CardTitle>Достижения</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex flex-wrap gap-2">
            <Badge variant="secondary">Новичок</Badge>
            {stats && stats.solved_riddles >= 5 && (
              <Badge variant="secondary">Решатель</Badge>
            )}
            {stats && stats.solved_riddles >= 20 && (
              <Badge variant="secondary">Эксперт</Badge>
            )}
            {stats && stats.success_rate >= 80 && (
              <Badge variant="secondary">Мастер</Badge>
            )}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}