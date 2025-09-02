"use client";

import { useState, useMemo } from "react";
import { useRiddles } from "@/hooks/useRiddles";
import { RiddleCard } from "@/components/riddles/RiddleCard";
import { Button } from "@/components/ui/button";

export default function RiddlesPage() {
  const { data: riddles, isLoading, error } = useRiddles();
  const [searchTerm, setSearchTerm] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("all");
  const [selectedDifficulty, setSelectedDifficulty] = useState("all");
  const [showFavoritesOnly, setShowFavoritesOnly] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 9;

  // Filter riddles based on search term, category, difficulty, and favorites
  const filteredRiddles = useMemo(() => {
    return riddles?.filter((riddle) => {
      // Search term filter
      const matchesSearch = riddle.riddle.title
        .toLowerCase()
        .includes(searchTerm.toLowerCase()) ||
        riddle.riddle.description
          .toLowerCase()
          .includes(searchTerm.toLowerCase());

      // Category filter
      const matchesCategory =
        selectedCategory === "all" ||
        riddle.riddle.category.name === selectedCategory;

      // Difficulty filter
      const matchesDifficulty =
        selectedDifficulty === "all" || riddle.riddle.difficulty === selectedDifficulty;

      // Favorites filter
      const matchesFavorites = !showFavoritesOnly || riddle.is_favorite;

      return matchesSearch && matchesCategory && matchesDifficulty && matchesFavorites;
    });
  }, [riddles, searchTerm, selectedCategory, selectedDifficulty, showFavoritesOnly]);

  // Pagination
  const totalPages = useMemo(() => {
    return filteredRiddles ? Math.ceil(filteredRiddles.length / itemsPerPage) : 0;
  }, [filteredRiddles]);

  const paginatedRiddles = useMemo(() => {
    if (!filteredRiddles) return [];
    const startIndex = (currentPage - 1) * itemsPerPage;
    return filteredRiddles.slice(startIndex, startIndex + itemsPerPage);
  }, [filteredRiddles, currentPage]);

  // Get unique categories for filter dropdown
  const categories = Array.from(
    new Set(riddles?.map((r) => r.riddle.category.name) || [])
  );

  // Show loading state
  if (isLoading) {
    return <div className="flex justify-center items-center h-64">Загрузка...</div>;
  }

  // Show error state
  if (error) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-red-500">Ошибка загрузки данных: {error.message}</div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-3xl font-bold">Все загадки</h1>

      {/* Filters */}
      <div className="bg-white p-4 rounded-lg shadow">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          {/* Search */}
          <div>
            <label htmlFor="search" className="block text-sm font-medium text-gray-700 mb-1">
              Поиск
            </label>
            <input
              type="text"
              id="search"
              placeholder="Введите текст..."
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>

          {/* Category */}
          <div>
            <label htmlFor="category" className="block text-sm font-medium text-gray-700 mb-1">
              Категория
            </label>
            <select
              id="category"
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
            >
              <option value="all">Все категории</option>
              {categories.map((category) => (
                <option key={category} value={category}>
                  {category}
                </option>
              ))}
            </select>
          </div>

          {/* Difficulty */}
          <div>
            <label htmlFor="difficulty" className="block text-sm font-medium text-gray-700 mb-1">
              Сложность
            </label>
            <select
              id="difficulty"
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedDifficulty}
              onChange={(e) => setSelectedDifficulty(e.target.value)}
            >
              <option value="all">Все уровни</option>
              <option value="easy">Легко</option>
              <option value="medium">Средне</option>
              <option value="hard">Сложно</option>
            </select>
          </div>

          {/* Favorites */}
          <div className="flex items-end">
            <label className="flex items-center">
              <input
                type="checkbox"
                className="rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                checked={showFavoritesOnly}
                onChange={(e) => setShowFavoritesOnly(e.target.checked)}
              />
              <span className="ml-2 text-sm text-gray-700">Только избранные</span>
            </label>
          </div>
        </div>
      </div>

      {/* Results count */}
      {filteredRiddles && (
        <div className="text-sm text-gray-600">
          Найдено загадок: {filteredRiddles.length}
        </div>
      )}

      {/* Riddles List */}
      {paginatedRiddles && paginatedRiddles.length > 0 ? (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {paginatedRiddles.map((riddle) => (
              <RiddleCard key={riddle.riddle.id} riddle={riddle} />
            ))}
          </div>

          {/* Pagination */}
          {totalPages > 1 && (
            <div className="flex justify-center items-center space-x-2">
              <Button
                variant="outline"
                onClick={() => setCurrentPage(prev => Math.max(prev - 1, 1))}
                disabled={currentPage === 1}
              >
                Назад
              </Button>
              
              <span className="text-sm">
                Страница {currentPage} из {totalPages}
              </span>
              
              <Button
                variant="outline"
                onClick={() => setCurrentPage(prev => Math.min(prev + 1, totalPages))}
                disabled={currentPage === totalPages}
              >
                Вперед
              </Button>
            </div>
          )}
        </>
      ) : (
        <div className="text-center py-8 text-gray-500">
          Загадки не найдены
        </div>
      )}
    </div>
  );
}