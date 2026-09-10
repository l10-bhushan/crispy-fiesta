import { ArrowRight, Eye, EyeOff, Lock, Mail, User } from "lucide-react";
import { useState } from "react";
import { Link } from "react-router-dom";

import Logo from "../components/Logo";
import Input from "../components/ui/Input";
import Button from "../components/ui/Button";

export default function Register() {
  const [showPassword, setShowPassword] = useState(false);

  return (
    <div className="dot-background flex min-h-screen flex-col">
      <header className="flex items-center justify-between px-6 py-6 md:px-10">
        <Logo />

        <p className="text-sm text-[#66736F]">
          Already have an account?{" "}
          <Link
            to="/login"
            className="font-semibold text-[#123C35] hover:underline"
          >
            Login
          </Link>
        </p>
      </header>

      <main className="flex flex-1 items-center justify-center px-6 py-10">
        <div className="w-full max-w-md rounded-3xl border border-[#E9E5D9] bg-white p-8 shadow-[0_20px_60px_rgba(30,50,45,0.06)] md:p-10">
          <div className="mb-8">
            <h1 className="text-2xl font-bold text-[#123C35]">
              Create your account
            </h1>

            <p className="mt-2 text-sm text-[#66736F]">
              Get started with Shawty in seconds
            </p>
          </div>

          <form className="space-y-5">
            <div className="relative">
              <Input label="Full name" type="text" placeholder="John Doe" />

              <User className="absolute right-4 top-11 h-4 w-4 text-[#9AA39F]" />
            </div>

            {/* <div className="relative">
              <Input label="Last name" type="text" placeholder="John Doe" />

              <User className="absolute right-4 top-[38px] h-4 w-4 text-[#9AA39F]" />
            </div>

            <div className="relative">
              <Input label="Username" type="text" placeholder="John Doe" />

              <User className="absolute right-4 top-[38px] h-4 w-4 text-[#9AA39F]" />
            </div> */}

            <div className="relative">
              <Input
                label="Email address"
                type="email"
                placeholder="you@example.com"
              />

              <Mail className="absolute right-4 top-11 h-4 w-4 text-[#9AA39F]" />
            </div>

            <div className="relative">
              <Input
                label="Password"
                type={showPassword ? "text" : "password"}
                placeholder="Create a strong password"
              />

              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-4 top-11 text-[#9AA39F]"
              >
                {showPassword ? (
                  <EyeOff className="h-4 w-4" />
                ) : (
                  <Eye className="h-4 w-4" />
                )}
              </button>
            </div>

            <Button type="submit" className="w-full">
              Create Account
              <ArrowRight className="ml-2 h-4 w-4" />
            </Button>
          </form>

          {/* <div className="my-7 flex items-center gap-4">
            <div className="h-px flex-1 bg-[#E9E5D9]" />
            <span className="text-xs text-[#A2AAA6]">OR</span>
            <div className="h-px flex-1 bg-[#E9E5D9]" />
          </div>

          <button className="flex w-full items-center justify-center gap-3 rounded-xl border border-[#DEDACD] bg-white py-3 text-sm font-medium text-[#17332F] transition hover:bg-[#FAF7ED]">
            <span className="font-bold text-[#4285F4]">G</span>
            Continue with Google
          </button> */}

          <p className="mt-6 text-center text-xs leading-5 text-[#8A9490]">
            By creating an account, you agree to our{" "}
            <a href="#" className="font-medium text-[#123C35]">
              Terms of Service
            </a>{" "}
            and{" "}
            <a href="#" className="font-medium text-[#123C35]">
              Privacy Policy
            </a>
            .
          </p>
        </div>
      </main>

      <footer className="px-6 py-6 text-center text-xs text-[#8A9490]">
        © {new Date().getFullYear()} Shawty. All rights reserved.
      </footer>
    </div>
  );
}
