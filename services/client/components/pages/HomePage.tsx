import Link from "next/link";
import Button from "../form/Button";
import Spacer from "../form/Spacer";

export default () => {
    async function logout () {
        const answer = await fetch("/api/v1/auth/logout", {
            method: "POST",
        })
    }

    return <div className={"w-screen min-h-screen flex justify-center items-center"}>
        <Spacer className={"m-2"}>
            <Link href={"/auth/sign-up"}>
                <Button text={"Sign Up"} />
            </Link>
        </Spacer>

        {/*<Spacer className={"m-2"}>*/}
        {/*    <Link href={"/auth/login"}>*/}
        {/*        <Button text={"Login"} />*/}
        {/*    </Link>*/}
        {/*</Spacer>*/}
        <br/>
    </div>
}
