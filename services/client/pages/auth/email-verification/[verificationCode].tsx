import {useEffect} from "react";
import {useRouter} from "next/router";

export default (props) => {
    const router = useRouter()

    const { verificationCode } = router.query

    useEffect(() => {
        (async () => {
            if (verificationCode) {
                const unencodedVerificationCode = decodeUrlSafeBase64(verificationCode);
                const answer = await fetch("/api/internal/v1/auth/email-verification/confirm", {
                    method: "PATCH",
                    body: JSON.stringify({ "verificationCode": unencodedVerificationCode })
                })
                const answerJson = await answer.json();
                if (answerJson.status !== "success") {
                    console.error(answerJson);
                    return;
                }
            }
        })()
    }, [verificationCode])

    return <>...</>
}

function decodeUrlSafeBase64(encoded) {
    encoded = encoded.replaceAll(/-/g, '+').replaceAll(/_/g, '/');
    return Buffer.from(encoded, 'base64').toString();
}
