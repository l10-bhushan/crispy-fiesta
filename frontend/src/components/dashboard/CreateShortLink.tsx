import { Link2, ArrowRight } from "lucide-react";
import { useState } from "react";

export default function CreateShortLink() {
  const [url, setUrl] = useState("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!url.trim()) {
      return;
    }

    console.log("Create short URL:", url);

    // Later:
    // POST /api/urls
  };

  return (
    <section className="rounded-2xl border border-[#E9E5D9] bg-white p-6 md:p-7">
      <div className="flex items-start gap-4">
        <div className="flex h-11 w-11 shrink-0 items-center justify-center rounded-2xl bg-[#FAF7ED]">
          <Link2 className="h-5 w-5 text-[#123C35]" />
        </div>

        <div>
          <h2 className="font-semibold text-[#17332F]">
            Create a new short link
          </h2>

          <p className="mt-1 text-sm text-[#66736F]">
            Paste your long URL below and get a short, shareable link instantly.
          </p>
        </div>
      </div>

      <form
        onSubmit={handleSubmit}
        className="mt-6 flex flex-col gap-3 md:flex-row"
      >
        <input
          type="url"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="https://example.com/very-long-url"
          className="
            min-w-0 flex-1 rounded-xl
            border border-[#DEDACD]
            bg-white px-4 py-3
            text-sm text-[#17332F]
            outline-none
            placeholder:text-[#9AA39F]
            focus:border-[#123C35]
            focus:ring-2
            focus:ring-[#123C35]/10
          "
        />

        <button
          type="submit"
          className="
            flex items-center justify-center gap-2
            rounded-xl bg-[#123C35]
            px-6 py-3
            text-sm font-semibold text-white
            transition hover:bg-[#1D5148]
          "
        >
          <Link2 className="h-4 w-4" />
          Shorten
          <ArrowRight className="h-4 w-4" />
        </button>
      </form>
    </section>
  );
}
