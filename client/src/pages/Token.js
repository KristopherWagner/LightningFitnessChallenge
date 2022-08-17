import { useSearchParams } from "react-router-dom";

export default function Token() {
    const [params, ] = useSearchParams();
    const tokenInfo = {
        code: params.get("code"),
        scope: params.get("scope"),
        state: params.get("state"),
    }
    return <p>{JSON.stringify(tokenInfo, 0, 2)}</p>
}