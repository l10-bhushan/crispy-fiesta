const topLinks = [
  { short: "shw.ty/7a3k", clicks: 124 },
  { short: "shw.ty/b9f2", clicks: 86 },
  { short: "shw.ty/k4m8", clicks: 52 },
  { short: "shw.ty/p6d1", clicks: 37 },
  { short: "shw.ty/t9z0", clicks: 28 },
];

export default function TopLinks() {
  return (
    <section className="rounded-2xl border border-[#E9E5D9] bg-white p-6">
      <div className="flex items-center justify-between">
        <h2 className="font-semibold text-[#17332F]">Top Links</h2>

        <button className="text-xs font-medium text-[#365F56]">
          View all →
        </button>
      </div>

      <div className="mt-5 space-y-4">
        {topLinks.map((link, index) => (
          <div key={link.short} className="flex items-center justify-between">
            <div className="flex min-w-0 items-center gap-3">
              <span className="w-4 text-xs text-[#8A9490]">{index + 1}</span>

              <span className="rounded-full bg-[#F2F5EF] px-3 py-1 text-xs font-semibold text-[#365F56]">
                {link.short}
              </span>
            </div>

            <span className="text-xs text-[#66736F]">{link.clicks} clicks</span>
          </div>
        ))}
      </div>
    </section>
  );
}
