import { useAddEnvironment, useGetAllEnvironments, useGetAllTitles } from "@/api/hooks";
import Spinner from "@/components/Spinner";
import { yupTitleEnvValidation } from "@/types/api/validation";
import { yupResolver } from "@hookform/resolvers/yup";
import { Alert, AlertTitle, Button, FormControl, InputLabel, MenuItem, Select, TextField } from "@mui/material";
import Box from "@mui/material/Box";
import { useEffect } from "react";
import { Controller, type SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";

interface AddEnvironmentForm {
    title: string;
    environment: string;
}

const validationSchemaAddEnvironment = yup.object({
    title: yupTitleEnvValidation.title,
    environment: yupTitleEnvValidation.environment,
});

export const AddEnvironmentForm: React.FC<{ addedTitle: string }> = ({ addedTitle }) => {
    const { mutate } = useGetAllEnvironments();
    const { data, isLoading, error } = useGetAllTitles();
    const { trigger, isMutating } = useAddEnvironment();

    const {
        register,
        handleSubmit,
        formState: { errors, isValid },
        control,
        setValue,
        resetField,
    } = useForm<AddEnvironmentForm>({
        resolver: yupResolver<AddEnvironmentForm>(validationSchemaAddEnvironment),
        mode: "onChange",
    });

    // When a title gets added, automatically set it as the selected value
    useEffect(() => {
        if (addedTitle) {
            setValue("title", addedTitle);
        }
    }, [addedTitle, setValue]);

    const onSubmit: SubmitHandler<AddEnvironmentForm> = formData => {
        trigger(formData).then(() => {
            resetField("environment");
            mutate();
        });
    };

    if (error) {
        return (
            <Alert severity="error">
                <AlertTitle>Could not load titles</AlertTitle>
                {error?.error}
            </Alert>
        );
    }
    if (isLoading) {
        return <Spinner />;
    }

    return (
        <Box
            component="form"
            noValidate
            onSubmit={handleSubmit(onSubmit)}
            sx={{
                display: "grid",
                columnGap: 2,
                gridTemplateColumns: "300px 300px 100px",
                alignItems: "start",
                maxWidth: 800,
            }}
        >
            <FormControl fullWidth>
                <InputLabel>Title</InputLabel>
                <Controller
                    name="title"
                    control={control}
                    defaultValue={data ? data[0]?.title : ""}
                    render={({ field }) => (
                        <Select {...field} size="small" label="Title">
                            {data?.map(t => (
                                <MenuItem key={t.title} value={t.title}>
                                    {t.title}
                                </MenuItem>
                            ))}
                        </Select>
                    )}
                />
            </FormControl>
            <TextField
                size="small"
                variant="outlined"
                required
                id="environment"
                label="Environment"
                error={Boolean(errors.environment)}
                helperText={errors.environment?.message}
                {...register("environment")}
            />
            <Button disabled={!isValid || isMutating} type="submit" fullWidth variant="contained">
                Add
            </Button>
        </Box>
    );
};
