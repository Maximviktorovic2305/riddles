"use client";

import { useDailyRiddle, useRiddles } from "@/hooks/useRiddles";
import { DailyRiddle } from "@/components/riddles/DailyRiddle";
import { RiddleCard } from "@/components/riddles/RiddleCard";

export default function Home() {
  const { data: dailyRiddle, isLoading: isDailyLoading, error: dailyError } = useDailyRiddle();
  const { data: riddles, isLoading: isRiddlesLoading, error: riddlesError } = useRiddles();

  // Show loading state
  if (isDailyLoading || isRiddlesLoading) {
    return <div className="flex justify-center items-center h-64">Загрузка...</div>;
  }

  // Show error state
  if (dailyError || riddlesError) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-red-500">Ошибка загрузки данных</div>
      </div>
    );
  }

  return (
    <div className="space-y-8">
      {/* Daily Riddle Section */}
      <section>
        <h2 className="text-2xl font-bold mb-4">Загадка дня</h2>
        {dailyRiddle ? (
          <DailyRiddle riddle={dailyRiddle} />
        ) : (
          <div className="text-center py-8 text-gray-500">
            Загадка дня пока не доступна
          </div>
        )}
      </section>

      {/* Random Riddles Section */}
      <section>
        <h2 className="text-2xl font-bold mb-4">Случайные загадки</h2>
        {riddles && riddles.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {riddles.slice(0, 6).map((riddle) => (
              <RiddleCard key={riddle.riddle.id} riddle={riddle} />
            ))}
          </div>
        ) : (
          <div className="text-center py-8 text-gray-500">
            Загадки не найдены
          </div>
        )}
      </section>
    </div>
  );
}
