import Logo from "@/components/shared/Logo"
import Navbar from "@/components/shared/Navbar"
import { NextPage } from "next"

const LoginPage: NextPage = () => {
  return (
    <>
      <Navbar>
        <Logo />
      </Navbar>

      <form className="max-w-3xl mx-auto mt-8 bg-white overflow-hidden border border-gray-300 rounded">
        <div
          style={{
            fontSize: "14.5px",
            backgroundColor: "rgba(0, 0, 0, 0.03)",
            padding: "0.75rem 1.25rem",
            borderBottom: "1px solid rgba(0, 0, 0, 0.125)",
          }}
        >
          ログインフォーム
        </div>

        <div className="py-8">
          <div className="flex px-3 mb-6 justify-center items-center">
            <label className="block w-1/6 text-gray-700 text-sm mb-2 text-right">
              ログインID
            </label>

            <div className="w-3/6 ml-6 mb-3">
              <input
                id="email"
                className="w-full appearance-none blocktext-gray-700 py-3 px-4 border-2 outline-none rounded leading-tight focus:border-gray-500 transition-all duration-200"
                type="email"
              />

              <p
                v-if="form.email.$dirty && form.email.$anyInvalid"
                className="emailErrorMessage w-full mt-2 text-red-500 text-xs transition-all duration-200"
              >
                {/* {{ "" }} */}
              </p>
            </div>
          </div>

          <div className="flex px-3 mb-6 justify-center items-center">
            <label className="block w-1/6 text-gray-700 text-sm mb-2 text-right">
              パスワード
            </label>

            <div className="w-3/6 ml-6 mb-3">
              <input
                id="password"
                className="w-full appearance-none blocktext-gray-700 py-3 px-4 border-2 outline-none rounded leading-tight focus:border-gray-500 transition-all duration-200"
                type="password"
              />

              <p className="passwordErrorMessage w-full mt-2 text-red-500 text-xs transition-all duration-200">
                {/* {{ form.password.required.$message }} */}
              </p>

              <p className="errorMessage w-full mt-2 text-red-500 text-xs transition-all duration-200">
                {/* {{ message }} */}
              </p>
            </div>
          </div>

          <div className="flex items-center justify-center">
            <button
              className="loginButton flex text-white py-2 px-4 rounded focus:outline-none focus:shadow-outline hover:opacity-80 transition-all duration-300"
              style={{ backgroundColor: "#1976D2" }}
              type="button"
            >
              ログイン
            </button>
          </div>
        </div>
      </form>
    </>
  )
}

export default LoginPage
