import Link from "next/link";
import Button from "../form/Button";

export default () => {
    async function logout () {
        const answer = await fetch("/api/v1/auth/logout", {
            method: "POST",
        })
    }

    return <>
        <h1 className={"text-3xl font-bold underline"}>Home</h1>
        <Link href={"/auth/sign-up"}>
            <Button text={"Sign Up"} />
        </Link>
        <Link href={"/auth/login"}>
            <Button text={"Login"} />
        </Link>
        <Button onClick={logout} loading={true} text={"Logout"} disabled={true} />
        <br/>
        <Button text={"Yoman"} fullWidth={true} />
    </>
}
