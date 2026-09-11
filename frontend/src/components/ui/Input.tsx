import type { InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string;
  error?: string;
}

export default function Input({
  label,
  error,
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
          w-full rounded-xl border
          bg-white px-4 py-3
          text-sm text-[#17332F]
          outline-none
          transition
          placeholder:text-[#A2AAA6]

          ${
            error
              ? "border-red-400 focus:border-red-500 focus:ring-red-500/10"
              : "border-[#DEDACD] focus:border-[#123C35] focus:ring-[#123C35]/10"
          }

          focus:ring-2
          ${className}
        `}
        {...props}
      />

      {error && <p className="text-xs font-medium text-red-600">{error}</p>}
    </div>
  );
}
