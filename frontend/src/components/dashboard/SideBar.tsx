import { BarChart3, ExternalLink, Home, Link2, Settings } from "lucide-react";
import { NavLink } from "react-router-dom";

import Logo from "../Logo";

const navigation = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: Home,
  },
  {
    name: "My Links",
    href: "/links",
    icon: Link2,
  },
  {
    name: "Analytics",
    href: "/analytics",
    icon: BarChart3,
  },
  {
    name: "Settings",
    href: "/settings",
    icon: Settings,
  },
];

export default function Sidebar() {
  return (
    <aside className="hidden w-64 shrink-0 border-r border-[#E9E5D9] bg-[#FFFDF7] lg:flex lg:flex-col">
      {/* Logo */}

      <div className="px-8 py-7">
        <Logo className="text-3xl" />
      </div>

      {/* Navigation */}

      <nav className="flex-1 px-4 pt-5">
        <div className="space-y-2">
          {navigation.map((item) => {
            const Icon = item.icon;

            return (
              <NavLink
                key={item.name}
                to={item.href}
                className={({ isActive }) =>
                  `
                  flex items-center gap-3 rounded-xl px-4 py-3
                  text-sm font-medium transition
                  ${
                    isActive
                      ? "bg-[#FAF7ED] text-[#123C35]"
                      : "text-[#66736F] hover:bg-[#FAF7ED] hover:text-[#123C35]"
                  }
                  `
                }
              >
                <Icon className="h-4.5 w-4.5" />

                {item.name}
              </NavLink>
            );
          })}
        </div>
      </nav>

      {/* Upgrade */}

      <div className="p-4">
        <div className="rounded-2xl border border-[#E9E5D9] bg-[#FAF7ED] p-5">
          <div className="mb-4 flex h-10 w-10 items-center justify-center rounded-xl bg-white">
            <Link2 className="h-5 w-5 text-[#123C35]" />
          </div>

          <h3 className="font-semibold text-[#123C35]">
            Shorter links.
            <br />
            Bigger possibilities.
          </h3>

          <p className="mt-2 text-xs leading-5 text-[#66736F]">
            Upgrade to unlock advanced analytics and more.
          </p>

          <button className="mt-4 flex items-center gap-2 rounded-lg bg-[#123C35] px-4 py-2.5 text-xs font-semibold text-white transition hover:bg-[#1D5148]">
            Upgrade to Pro
            <ExternalLink className="h-3.5 w-3.5" />
          </button>
        </div>
      </div>
    </aside>
  );
}
