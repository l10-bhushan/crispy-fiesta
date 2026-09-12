import { ArrowRight, Link2 } from "lucide-react";

export default function UpgradeCard() {
  return (
    <section className="rounded-2xl border border-[#E9E5D9] bg-[#FAF7ED] p-6">
      <div className="flex items-start gap-4">
        <div className="flex h-11 w-11 shrink-0 items-center justify-center rounded-xl bg-white">
          <Link2 className="h-5 w-5 text-[#123C35]" />
        </div>

        <div>
          <h2 className="font-semibold text-[#17332F]">
            Want more from Shawty?
          </h2>

          <p className="mt-2 text-sm leading-6 text-[#66736F]">
            Upgrade to get advanced analytics, custom domains and more.
          </p>

          <button className="mt-4 flex items-center gap-2 rounded-xl bg-[#123C35] px-4 py-2.5 text-xs font-semibold text-white hover:bg-[#1D5148]">
            Upgrade to Pro
            <ArrowRight className="h-3.5 w-3.5" />
          </button>
        </div>
      </div>
    </section>
  );
}
