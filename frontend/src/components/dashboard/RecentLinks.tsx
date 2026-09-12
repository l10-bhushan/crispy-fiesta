import { Copy, ExternalLink, Link2, MoreVertical } from "lucide-react";

const links = [
  {
    short: "shw.ty/7a3k",
    original: "https://www.example.com/product/...",
    clicks: 124,
    created: "2 hours ago",
  },
  {
    short: "shw.ty/b9f2",
    original: "https://github.com/your-repo",
    clicks: 86,
    created: "5 hours ago",
  },
  {
    short: "shw.ty/k4m8",
    original: "https://medium.com/interesting-article",
    clicks: 52,
    created: "1 day ago",
  },
  {
    short: "shw.ty/p6d1",
    original: "https://docs.google.com/spreadsheets/...",
    clicks: 37,
    created: "1 day ago",
  },
  {
    short: "shw.ty/t9z0",
    original: "https://www.youtube.com/watch?v=...",
    clicks: 28,
    created: "2 days ago",
  },
  {
    short: "shw.ty/x8v3",
    original: "https://www.figma.com/file/...",
    clicks: 21,
    created: "3 days ago",
  },
];

export default function RecentLinks() {
  return (
    <section className="rounded-2xl border border-[#E9E5D9] bg-white">
      <div className="flex items-center justify-between border-b border-[#E9E5D9] px-6 py-5">
        <div className="flex items-center gap-3">
          <Link2 className="h-5 w-5 text-[#123C35]" />

          <h2 className="font-semibold text-[#17332F]">Recent Links</h2>
        </div>

        <button className="text-sm font-medium text-[#365F56] hover:text-[#123C35]">
          View all →
        </button>
      </div>

      {/* Desktop table */}

      <div className="hidden overflow-x-auto md:block">
        <table className="w-full">
          <thead>
            <tr className="border-b border-[#E9E5D9] text-left">
              <th className="px-5 py-4 text-xs font-medium text-[#8A9490]">
                Short Link
              </th>

              <th className="px-5 py-4 text-xs font-medium text-[#8A9490]">
                Original URL
              </th>

              <th className="px-5 py-4 text-xs font-medium text-[#8A9490]">
                Clicks
              </th>

              <th className="px-5 py-4 text-xs font-medium text-[#8A9490]">
                Created
              </th>

              <th className="px-5 py-4 text-xs font-medium text-[#8A9490]">
                Actions
              </th>
            </tr>
          </thead>

          <tbody>
            {links.map((link) => (
              <tr
                key={link.short}
                className="border-b border-[#E9E5D9] last:border-0"
              >
                <td className="px-5 py-5">
                  <div className="flex items-center gap-2">
                    <Link2 className="h-4 w-4 text-[#123C35]" />

                    <span className="rounded-full bg-[#F2F5EF] px-3 py-1 text-xs font-semibold text-[#365F56]">
                      {link.short}
                    </span>
                  </div>
                </td>

                <td className="max-w-62.5 truncate px-5 py-5 text-sm text-[#66736F]">
                  {link.original}
                </td>

                <td className="px-5 py-5 text-sm font-medium text-[#17332F]">
                  {link.clicks}
                </td>

                <td className="whitespace-nowrap px-5 py-5 text-sm text-[#66736F]">
                  {link.created}
                </td>

                <td className="px-5 py-5">
                  <div className="flex items-center gap-3">
                    <button
                      title="Copy"
                      className="text-[#66736F] hover:text-[#123C35]"
                    >
                      <Copy className="h-4 w-4" />
                    </button>

                    <button
                      title="Open"
                      className="text-[#66736F] hover:text-[#123C35]"
                    >
                      <ExternalLink className="h-4 w-4" />
                    </button>

                    <button className="text-[#66736F] hover:text-[#123C35]">
                      <MoreVertical className="h-4 w-4" />
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Mobile */}

      <div className="divide-y divide-[#E9E5D9] md:hidden">
        {links.map((link) => (
          <div key={link.short} className="p-5">
            <div className="flex items-center justify-between">
              <span className="rounded-full bg-[#F2F5EF] px-3 py-1 text-xs font-semibold text-[#365F56]">
                {link.short}
              </span>

              <span className="text-xs text-[#8A9490]">{link.created}</span>
            </div>

            <p className="mt-3 truncate text-sm text-[#66736F]">
              {link.original}
            </p>

            <div className="mt-3 flex items-center justify-between">
              <span className="text-xs text-[#66736F]">
                {link.clicks} clicks
              </span>

              <div className="flex gap-3">
                <Copy className="h-4 w-4 text-[#66736F]" />
                <ExternalLink className="h-4 w-4 text-[#66736F]" />
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
}
