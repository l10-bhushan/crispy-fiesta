import type { LucideIcon } from "lucide-react";

interface StatCardProps {
  title: string;
  value: string;
  change: string;
  icon: LucideIcon;
}

export default function StatCard({
  title,
  value,
  change,
  icon: Icon,
}: StatCardProps) {
  return (
    <div className="rounded-2xl border border-[#E9E5D9] bg-white p-5">
      <div className="flex items-center justify-between">
        <p className="text-sm text-[#66736F]">{title}</p>

        <div className="flex h-9 w-9 items-center justify-center rounded-xl bg-[#FAF7ED]">
          <Icon className="h-4.5 w-4.5 text-[#123C35]" />
        </div>
      </div>

      <p className="mt-4 text-2xl font-bold tracking-tight text-[#123C35]">
        {value}
      </p>

      <p className="mt-2 text-xs font-medium text-[#3E806A]">↑ {change}</p>
    </div>
  );
}
