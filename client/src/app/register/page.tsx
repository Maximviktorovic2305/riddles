"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { RegisterForm } from "@/components/auth/RegisterForm";
import Link from "next/link";

export default function RegisterPage() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl font-bold text-center">Регистрация</CardTitle>
          <CardDescription className="text-center">
            Создайте новую учетную запись
          </CardDescription>
        </CardHeader>
        <CardContent>
          <RegisterForm />
          <div className="mt-4 text-center text-sm">
            Уже есть аккаунт?{" "}
            <Link href="/login" className="text-blue-600 hover:underline">
              Войдите
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}