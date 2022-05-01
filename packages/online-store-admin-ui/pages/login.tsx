import { api } from "@/api"
import Logo from "@/components/shared/Logo"
import Navbar from "@/components/shared/Navbar"
import { NextPage } from "next"
import { useState } from "react"

const LoginPage: NextPage = () => {
  const [loginId, setLoginId] = useState("")
  const [password, setPassword] = useState("")

  const login = async () => {
    const response = await api.postAdminLogin({ loginId, password })
    console.log(response.token)
  }

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
                className="w-full appearance-none blocktext-gray-700 py-3 px-4 border-2 outline-none rounded leading-tight focus:border-gray-500 transition-all duration-200"
                type="text"
                value={loginId}
                onChange={(e) => setLoginId(e.target.value)}
              />
            </div>
          </div>

          <div className="flex px-3 mb-6 justify-center items-center">
            <label className="block w-1/6 text-gray-700 text-sm mb-2 text-right">
              パスワード
            </label>

            <div className="w-3/6 ml-6 mb-3">
              <input
                className="w-full appearance-none blocktext-gray-700 py-3 px-4 border-2 outline-none rounded leading-tight focus:border-gray-500 transition-all duration-200"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
          </div>

          <div className="flex items-center justify-center">
            <button
              className="loginButton flex text-white py-2 px-4 rounded focus:outline-none focus:shadow-outline hover:opacity-80 transition-all duration-300"
              style={{ backgroundColor: "#1976D2" }}
              type="button"
              onClick={login}
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
