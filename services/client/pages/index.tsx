import Link from "next/link";

export default () => {
    return <>
        <h1>Home</h1>
        <Link href={"/auth/login"}>
            <button>
                Login
            </button>
        </Link
    ></>
}
