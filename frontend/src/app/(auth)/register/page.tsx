"use client";

import Link from "next/link";
import { useState } from "react";
import { useRegister } from "@/hooks/useAuth";

export default function RegisterPage() {
  const [form, setForm] = useState({ name: "", email: "", password: "" });
  const { mutate: register, isPending } = useRegister();

  const handleSubmit = () => {
    if (!form.name || !form.email || !form.password) return;
    register(form);
  };

  return (
    <div className="page-centered">
      <div className="auth-card">
        <div className="auth-header">
          <h1>Create account</h1>
          <p>Start tracking weather in your cities</p>
        </div>

        <div className="input-group">
          <label className="input-label">Name</label>
          <input
            className="input"
            type="text"
            placeholder="Emmanuel"
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.target.value })}
            onKeyDown={(e) => e.key === "Enter" && handleSubmit()}
          />
        </div>

        <div className="input-group">
          <label className="input-label">Email</label>
          <input
            className="input"
            type="email"
            placeholder="you@example.com"
            value={form.email}
            onChange={(e) => setForm({ ...form, email: e.target.value })}
            onKeyDown={(e) => e.key === "Enter" && handleSubmit()}
          />
        </div>

        <div className="input-group">
          <label className="input-label">Password</label>
          <input
            className="input"
            type="password"
            placeholder="••••••••"
            value={form.password}
            onChange={(e) => setForm({ ...form, password: e.target.value })}
            onKeyDown={(e) => e.key === "Enter" && handleSubmit()}
          />
        </div>

        <button
          className="btn btn-primary btn-full"
          onClick={handleSubmit}
          disabled={isPending}
        >
          {isPending ? "Creating account..." : "Create account"}
        </button>

        <div className="auth-footer">
          Already have an account? <Link href="/login">Sign in</Link>
        </div>
      </div>
    </div>
  );
}
