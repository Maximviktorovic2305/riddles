'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useAuth } from '@/hooks/useAuth';
import { Button } from '@/components/ui/button';

export function Navigation() {
  const pathname = usePathname();
  const { isAuthenticated, logout } = useAuth();

  const navigationItems = [
    { name: 'Главная', href: '/' },
    { name: 'Все загадки', href: '/riddles' },
  ];

  return (
    <nav className="bg-white shadow">
      <div className="container mx-auto px-4">
        <div className="flex justify-between h-16">
          <div className="flex">
            <div className="flex-shrink-0 flex items-center">
              <Link href="/" className="text-xl font-bold text-gray-800">
                Загадки
              </Link>
            </div>
            <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
              {navigationItems.map((item) => (
                <Link
                  key={item.name}
                  href={item.href}
                  className={`${
                    pathname === item.href
                      ? 'border-b-2 border-blue-500 text-gray-900'
                      : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700'
                  } inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium`}
                >
                  {item.name}
                </Link>
              ))}
            </div>
          </div>
          <div className="flex items-center">
            {isAuthenticated ? (
              <div className="flex items-center space-x-4">
                <Link href="/profile">
                  <Button variant="ghost">Профиль</Button>
                </Link>
                <Button variant="outline" onClick={logout}>
                  Выйти
                </Button>
              </div>
            ) : (
              <div className="flex space-x-2">
                <Link href="/login">
                  <Button variant="ghost">Войти</Button>
                </Link>
                <Link href="/register">
                  <Button variant="outline">Регистрация</Button>
                </Link>
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}