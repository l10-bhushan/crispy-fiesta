import { ChevronDown, Menu } from "lucide-react";

export default function DashboardHeader() {
  return (
    <header className="flex items-center justify-between border-b border-[#E9E5D9] px-6 py-5 md:px-8">
      <button className="rounded-lg p-2 text-[#66736F] hover:bg-[#FAF7ED] lg:hidden">
        <Menu className="h-5 w-5" />
      </button>

      <div className="ml-auto flex items-center gap-3">
        <div className="flex h-9 w-9 items-center justify-center rounded-full bg-[#FAF7ED] text-sm font-semibold text-[#123C35]">
          B
        </div>

        <span className="hidden text-sm font-medium text-[#17332F] sm:block">
          Bhushan Rajput
        </span>

        <ChevronDown className="h-4 w-4 text-[#66736F]" />
      </div>
    </header>
  );
}
