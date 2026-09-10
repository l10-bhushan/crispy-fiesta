import type { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string;
}

export default function Input({
  label,
  className = "",
  ...props
}: InputProps) {
  return (
    <div className="space-y-2">
      <label className="block text-sm font-medium text-[#17332F]">
        {label}
      </label>

      <input
        className={`
          w-full rounded-xl border border-[#DEDACD]
          bg-white px-4 py-3
          text-sm text-[#17332F]
          outline-none
          transition
          placeholder:text-[#A2AAA6]
          focus:border-[#123C35]
          focus:ring-2
          focus:ring-[#123C35]/10
          ${className}
        `}
        {...props}
      />
    </div>
  );
}