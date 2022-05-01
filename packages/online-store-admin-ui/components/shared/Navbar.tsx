import { FC, ReactNode } from "react"

const Navbar: FC<{ children?: ReactNode | undefined }> = ({ children }) => {
  return (
    <nav className="flex bg-white border-b border-gray-200 top-0 inset-x-0 z-100 h-16 items-center">
      <div className="px-12 w-full flex justify-between">{children}</div>
    </nav>
  )
}

export default Navbar
