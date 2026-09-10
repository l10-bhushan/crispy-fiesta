import { Link } from "react-router-dom";
import Logo from "../Logo";

export default function Header() {
  return (
    <header className="flex items-center justify-between px-6 py-5 md:px-10">
      <Logo />

      <div className="flex items-center gap-3">
        <Link
          to="/login"
          className="rounded-xl border border-[#DEDACD] bg-white px-5 py-2.5 text-sm font-medium text-[#17332F] transition hover:bg-[#FAF7ED]"
        >
          Login
        </Link>

        <Link
          to="/register"
          className="rounded-xl bg-[#123C35] px-5 py-2.5 text-sm font-semibold text-white transition hover:bg-[#1D5148]"
        >
          Register
        </Link>
      </div>
    </header>
  );
}
