"use client";

import { useUserPasswordResetRequest } from "@/api/hooks";
import { yupAuthValidation } from "@/types/api/validation";
import { yupResolver } from "@hookform/resolvers/yup";
import { Box, Button, Link, TextField, Typography } from "@mui/material";
import type React from "react";
import { useState } from "react";
import { type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";

interface FormInputPasswordForgot {
    email: string;
}

const validationSchemaPasswordForgot = yup.object({
    email: yupAuthValidation.email,
});

const PasswordForgot: React.FC = () => {
    const { trigger, isMutating, error } = useUserPasswordResetRequest();
    const [isSent, setIsSent] = useState(false);

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
    } = useForm<FormInputPasswordForgot>({
        resolver: yupResolver(validationSchemaPasswordForgot),
        mode: "onChange",
    });

    const onSubmit: SubmitHandler<FormInputPasswordForgot> = async data => {
        await trigger(data);
        setIsSent(true);
    };

    if (isSent && !error && !isMutating) {
        return <Typography align="center">Please check your email to set a new password</Typography>;
    }

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
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
                Request password reset
            </Button>
            <Link href="/auth/login" variant="body2" color="textSecondary" display="block" align="center">
                Log in instead?
            </Link>
        </Box>
    );
};

export default PasswordForgot;
