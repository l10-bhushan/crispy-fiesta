import {
  ArrowRight,
  BarChart3,
  Check,
  Link2,
  ShieldCheck,
  Zap,
} from "lucide-react";

import { Link } from "react-router-dom";

import Header from "../components/layout/Header";
import Footer from "../components/layout/Footer";
import Button from "../components/ui/Button";

const features = [
  {
    icon: Link2,
    title: "Shorten Links",
    description:
      "Turn long, messy URLs into clean, shareable links in seconds.",
  },
  {
    icon: BarChart3,
    title: "Track Performance",
    description:
      "See how many people click your links, and when they get clicked.",
  },
  {
    icon: ShieldCheck,
    title: "Built for Reliability",
    description: "Fast, secure and always available when you need your links.",
  },
  {
    icon: Zap,
    title: "Easy to Use",
    description: "No complicated setup. Just paste, shorten and share.",
  },
];

export default function Home() {
  return (
    <div className="dot-background min-h-screen">
      <Header />

      <main>
        {/* Hero */}

        <section className="mx-auto max-w-7xl px-6 py-20 md:px-10 md:py-28">
          <div className="grid items-center gap-16 lg:grid-cols-2">
            <div>
              <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-[#DEDACD] bg-white px-4 py-2 text-sm text-[#36504A]">
                <span className="h-2 w-2 rounded-full bg-[#123C35]" />
                Shorten · Share · Track
              </div>

              <h1 className="max-w-2xl text-5xl font-bold leading-[1.05] tracking-tight text-[#123C35] md:text-7xl">
                Shorter links.
                <br />
                Bigger possibilities.
              </h1>

              <p className="mt-7 max-w-xl text-lg leading-8 text-[#66736F]">
                Shawty is a simple and fast URL shortener. Turn long links into
                clean, short URLs and share them with confidence.
              </p>

              <div className="mt-9 flex flex-wrap gap-3">
                <Link to="/register">
                  <Button>
                    Get Started Free
                    <ArrowRight className="ml-2 h-4 w-4" />
                  </Button>
                </Link>

                <Button variant="secondary">Learn More</Button>
              </div>
            </div>

            {/* URL Shortener Preview */}

            <div className="relative flex justify-center">
              <div className="absolute h-80 w-80 rounded-full bg-[#F2EEDF] blur-3xl" />

              <div className="relative w-full max-w-md rounded-3xl border border-[#E9E5D9] bg-white p-6 shadow-[0_20px_60px_rgba(30,50,45,0.08)]">
                <div className="mb-6 flex items-center justify-between">
                  <span className="font-[Pacifico] text-xl text-[#123C35]">
                    Shawty
                  </span>

                  <Link2 className="h-5 w-5 text-[#123C35]" />
                </div>

                <div className="flex gap-2">
                  <input
                    readOnly
                    value="https://example.com/very-long-url"
                    className="min-w-0 flex-1 rounded-xl border border-[#DEDACD] bg-[#FAF7ED] px-4 py-3 text-sm text-[#66736F] outline-none"
                  />

                  <button className="rounded-xl bg-[#123C35] px-5 text-sm font-semibold text-white">
                    Shorten
                  </button>
                </div>

                <div className="mt-5 rounded-2xl border border-[#E9E5D9] bg-[#FAF7ED] p-4">
                  <div className="flex items-center gap-3">
                    <div className="flex h-9 w-9 items-center justify-center rounded-full bg-[#123C35]">
                      <Check className="h-4 w-4 text-white" />
                    </div>

                    <div className="flex-1">
                      <p className="text-sm font-semibold text-[#17332F]">
                        shw.ty/7a3k
                      </p>

                      <p className="mt-1 text-xs text-[#66736F]">
                        Your shortened URL is ready
                      </p>
                    </div>

                    <Link2 className="h-4 w-4 text-[#66736F]" />
                  </div>
                </div>

                <div className="mt-6 grid grid-cols-3 gap-4 border-t border-[#E9E5D9] pt-5">
                  <div>
                    <p className="text-lg font-bold text-[#123C35]">1.2k</p>
                    <p className="text-xs text-[#8A9490]">Clicks</p>
                  </div>

                  <div>
                    <p className="text-lg font-bold text-[#123C35]">24h</p>
                    <p className="text-xs text-[#8A9490]">Recent</p>
                  </div>

                  <div>
                    <p className="text-lg font-bold text-[#123C35]">98%</p>
                    <p className="text-xs text-[#8A9490]">Uptime</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Features */}

        <section className="border-t border-[#E9E5D9] bg-[#FFFDF7] px-6 py-20 md:px-10">
          <div className="mx-auto max-w-6xl">
            <div className="mx-auto max-w-2xl text-center">
              <span className="inline-flex rounded-full border border-[#DEDACD] bg-white px-4 py-2 text-xs font-medium text-[#36504A]">
                Why Shawty?
              </span>

              <h2 className="mt-5 text-3xl font-bold tracking-tight text-[#123C35] md:text-4xl">
                Simple. Fast. Reliable.
              </h2>

              <p className="mt-4 text-[#66736F]">
                Everything you need to create, manage and track your short
                links.
              </p>
            </div>

            <div className="mt-14 grid gap-10 md:grid-cols-2 lg:grid-cols-4">
              {features.map((feature) => {
                const Icon = feature.icon;

                return (
                  <div key={feature.title} className="text-center">
                    <div className="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-[#FAF7ED]">
                      <Icon className="h-6 w-6 text-[#123C35]" />
                    </div>

                    <h3 className="mt-5 font-semibold text-[#17332F]">
                      {feature.title}
                    </h3>

                    <p className="mt-2 text-sm leading-6 text-[#66736F]">
                      {feature.description}
                    </p>
                  </div>
                );
              })}
            </div>
          </div>
        </section>

        {/* CTA */}

        <section className="px-6 pb-20 md:px-10">
          <div className="mx-auto max-w-6xl overflow-hidden rounded-3xl bg-[#FAF7ED] px-8 py-12 md:px-12 md:py-14">
            <div className="flex flex-col items-start justify-between gap-8 md:flex-row md:items-center">
              <div>
                <span className="rounded-full border border-[#DEDACD] bg-white px-3 py-1 text-xs text-[#66736F]">
                  Ready to get started?
                </span>

                <h2 className="mt-4 max-w-lg text-3xl font-bold leading-tight text-[#123C35]">
                  Create your first short link in seconds.
                </h2>

                <p className="mt-3 max-w-lg text-[#66736F]">
                  Join Shawty and make your links shorter, cleaner and easier to
                  share.
                </p>
              </div>

              <Link to="/register">
                <Button>
                  Get Started Free
                  <ArrowRight className="ml-2 h-4 w-4" />
                </Button>
              </Link>
            </div>
          </div>
        </section>
      </main>

      <Footer />
    </div>
  );
}
