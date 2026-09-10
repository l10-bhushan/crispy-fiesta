import Logo from "../Logo";

export default function Footer() {
  return (
    <footer className="border-t border-[#E9E5D9] px-6 py-6 md:px-10">
      <div className="flex flex-col items-center justify-between gap-4 text-sm text-[#66736F] md:flex-row">
        <div className="flex items-center gap-4">
          <Logo className="text-lg" />

          <span className="hidden h-4 w-px bg-[#DEDACD] sm:block" />

          <span>Shorten links. Share more.</span>
        </div>

        <div className="flex items-center gap-5">
          <a href="#" className="hover:text-[#123C35]">
            About
          </a>

          <a href="#" className="hover:text-[#123C35]">
            Privacy
          </a>

          <a href="#" className="hover:text-[#123C35]">
            Terms
          </a>
        </div>
      </div>
    </footer>
  );
}
