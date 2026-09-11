import type { ButtonHTMLAttributes } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary";
}

export default function Button({
  variant = "primary",
  className = "",
  children,
  ...props
}: ButtonProps) {
  const base =
    "inline-flex items-center justify-center rounded-xl px-5 py-3 text-sm font-semibold transition-all duration-200 cursor-pointer";

  const variants = {
    primary:
      "bg-[#123C35] text-white hover:bg-[#1D5148] shadow-sm hover:shadow-md",
    secondary:
      "border border-[#DEDACD] bg-white text-[#17332F] hover:bg-[#FAF7ED]",
  };

  return (
    <button
      className={`${base} ${variants[variant]} ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}
