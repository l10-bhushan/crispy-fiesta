import { ChevronDown } from "lucide-react";

const points = [
  [10, 100],
  [65, 80],
  [120, 96],
  [175, 78],
  [230, 42],
  [285, 62],
  [340, 50],
];

export default function LinkPerformance() {
  const path = points
    .map(([x, y], index) => `${index === 0 ? "M" : "L"} ${x} ${y}`)
    .join(" ");

  return (
    <section className="rounded-2xl border border-[#E9E5D9] bg-white p-6">
      <div className="flex items-center justify-between">
        <h2 className="font-semibold text-[#17332F]">Link Performance</h2>

        <button className="flex items-center gap-1 text-xs text-[#66736F]">
          Last 7 days
          <ChevronDown className="h-3.5 w-3.5" />
        </button>
      </div>

      <div className="mt-6">
        <svg viewBox="0 0 360 150" className="h-auto w-full overflow-visible">
          {/* Grid */}

          {[20, 55, 90, 125].map((y) => (
            <line
              key={y}
              x1="10"
              y1={y}
              x2="350"
              y2={y}
              stroke="#E9E5D9"
              strokeWidth="1"
            />
          ))}

          {/* Area */}

          <path
            d={`${path} L 340 125 L 10 125 Z`}
            fill="rgba(18, 60, 53, 0.06)"
          />

          {/* Line */}

          <path
            d={path}
            fill="none"
            stroke="#123C35"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />

          {/* Points */}

          {points.map(([x, y]) => (
            <circle key={`${x}-${y}`} cx={x} cy={y} r="3" fill="#123C35" />
          ))}
        </svg>

        <div className="mt-2 flex justify-between text-[10px] text-[#8A9490]">
          <span>Sep 4</span>
          <span>Sep 5</span>
          <span>Sep 6</span>
          <span>Sep 7</span>
          <span>Sep 8</span>
          <span>Sep 9</span>
          <span>Sep 10</span>
        </div>
      </div>
    </section>
  );
}
