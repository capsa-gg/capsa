"use client";

import { useUserLogin } from "@/api/hooks";
import useUser from "@/context/UserContext";
import { removeJwtCookie, setJwtCookie } from "@/data/jwt/cookiesClient";
import { removeJwtFromLocalStorage, setJwtInLocalStorage } from "@/data/jwt/localStorage";
import { yupAuthValidation } from "@/types/api/validation";
import { yupResolver } from "@hookform/resolvers/yup";
import { Box, Button, Link, TextField } from "@mui/material";
import { useRouter, useSearchParams } from "next/navigation";
import type React from "react";
import { Suspense, useEffect } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";

interface FormInputLogin {
    email: string;
    password: string;
}

const validationSchemaLogin = yup.object({
    email: yupAuthValidation.email,
    password: yupAuthValidation.password,
});

const LoginPage: React.FC = () => {
    const { trigger, isMutating, error, data } = useUserLogin();
    const router = useRouter();
    const searchParams = useSearchParams();
    const user = useUser();
    const redirectUrl = searchParams.get("redirect");

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
    } = useForm<FormInputLogin>({
        resolver: yupResolver<FormInputLogin>(validationSchemaLogin),
        mode: "onChange",
    });

    // Remove jwt info if this page is reached
    // biome-ignore lint/correctness/useExhaustiveDependencies: fine here, for now
    useEffect(() => {
        if (redirectUrl) {
            removeJwtCookie();
            removeJwtFromLocalStorage();
        }
    }, []); // eslint-disable-line react-hooks/exhaustive-deps

    // When we have the response, set values, redirect
    // biome-ignore lint/correctness/useExhaustiveDependencies: fine here, for now
    useEffect(() => {
        if (!error && data) {
            setJwtCookie(data);
            setJwtInLocalStorage(data);
            user.reloadInfoFromLocalStorage(); // This will update the signed in status in nav

            if (redirectUrl) {
                router.replace(redirectUrl);
            } else {
                router.refresh();
            }
        }
    }, [data]); // eslint-disable-line react-hooks/exhaustive-deps

    const onSubmit: SubmitHandler<FormInputLogin> = formData => {
        trigger(formData);
    };

    return (
        <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
            <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email"
                autoComplete="email"
                autoFocus
                {...register("email")}
                error={Boolean(errors.email)}
                helperText={errors.email?.message}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                id="password"
                label="Password"
                type="password"
                autoComplete="current-password"
                {...register("password")}
                error={Boolean(errors.password)}
                helperText={errors.password?.message}
            />
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
                Login
            </Button>
            <Link href="/auth/password-forgot" variant="body2" color="textSecondary" display="block" align="center">
                Forgot your password?
            </Link>
        </Box>
    );
};

const SuspenseLoginPage = () => (
    <Suspense>
        <LoginPage />
    </Suspense>
);

export default SuspenseLoginPage;
