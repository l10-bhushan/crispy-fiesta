import { Link } from "react-router-dom";

interface LogoProps {
  className?: string;
}

export default function Logo({ className = "" }: LogoProps) {
  return (
    <Link
      to="/"
      className={`font-[Pacifico] text-2xl text-[#123C35] ${className}`}
    >
      Shawty
    </Link>
  );
}
