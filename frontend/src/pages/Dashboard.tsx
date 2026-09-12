import { BarChart3, Clock3, Link2, Users } from "lucide-react";

import Sidebar from "../components/dashboard/SideBar";
import DashboardHeader from "../components/dashboard/DashBoardHeader";
import StatCard from "../components/dashboard/StatCard";
import CreateShortLink from "../components/dashboard/CreateShortLink";
import RecentLinks from "../components/dashboard/RecentLinks";
import LinkPerformance from "../components/dashboard/PerformanceChart";
import TopLinks from "../components/dashboard/TopLinks";
import UpgradeCard from "../components/dashboard/UpgradeCard";
import Footer from "../components/layout/Footer";

export default function Dashboard() {
  return (
    <div className="dot-background flex min-h-screen">
      <Sidebar />

      <div className="flex min-w-0 flex-1 flex-col">
        <DashboardHeader />

        <main className="flex-1 px-6 py-8 md:px-8 lg:px-10">
          <div className="mx-auto max-w-350">
            {/* Heading */}

            <div className="mb-8">
              <p className="text-sm text-[#66736F]">Good morning, Bhushan 👋</p>

              <h1 className="mt-2 max-w-2xl text-3xl font-bold tracking-tight text-[#123C35] md:text-4xl">
                Turn your long links
                <br className="hidden md:block" />
                into something short.
              </h1>

              <p className="mt-3 text-sm text-[#66736F]">
                Create, manage and track your shortened links all in one place.
              </p>
            </div>

            {/* Stats */}

            <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
              <StatCard
                title="Total Links"
                value="12"
                change="2 this week"
                icon={Link2}
              />

              <StatCard
                title="Total Clicks"
                value="348"
                change="12% this week"
                icon={BarChart3}
              />

              <StatCard
                title="Unique Visitors"
                value="276"
                change="15% this week"
                icon={Users}
              />

              <StatCard
                title="Uptime"
                value="99.9%"
                change="0.1% this week"
                icon={Clock3}
              />
            </div>

            {/* Main content */}

            <div className="mt-6 grid gap-6 xl:grid-cols-[minmax(0,1fr)_375px]">
              {/* Left */}

              <div className="min-w-0 space-y-6">
                <CreateShortLink />

                <RecentLinks />
              </div>

              {/* Right */}

              <div className="space-y-6">
                <LinkPerformance />

                <TopLinks />

                <UpgradeCard />
              </div>
            </div>
          </div>
        </main>

        <Footer />
      </div>
    </div>
  );
}
