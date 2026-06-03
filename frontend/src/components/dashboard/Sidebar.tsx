"use client";

import { Cloud, Star, LogOut } from "lucide-react";
import { useLogout } from "@/hooks/useAuth";
import { getUser } from "@/lib/auth";

export default function Sidebar() {
  const logout = useLogout();
  const user = getUser();

  return (
    <aside className="sidebar">
      <div className="sidebar-logo">
        <Cloud size={22} color="var(--accent)" />
        <span>Nimbus</span>
      </div>

      <p className="section-label" style={{ paddingLeft: "0.875rem" }}>
        Menu
      </p>

      <button className="sidebar-nav-item active">
        <Cloud size={16} />
        Weather
      </button>

      <div className="sidebar-footer">
        <div style={{ padding: "0.5rem 0.875rem", marginBottom: "0.5rem" }}>
          <p
            style={{
              fontSize: "0.8rem",
              color: "var(--text-secondary)",
              fontWeight: 500,
            }}
          >
            {user?.name}
          </p>
        </div>
        <button className="sidebar-nav-item" onClick={logout}>
          <LogOut size={16} />
          Sign out
        </button>
      </div>
    </aside>
  );
}
