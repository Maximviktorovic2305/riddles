"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { LoginForm } from "@/components/auth/LoginForm";
import Link from "next/link";

export default function LoginPage() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl font-bold text-center">Вход</CardTitle>
          <CardDescription className="text-center">
            Войдите в свою учетную запись
          </CardDescription>
        </CardHeader>
        <CardContent>
          <LoginForm />
          <div className="mt-4 text-center text-sm">
            Нет аккаунта?{" "}
            <Link href="/register" className="text-blue-600 hover:underline">
              Зарегистрируйтесь
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}