#include <iostream>
#include <unordered_map>

std::string randomBase64 = "Nz746LU-BCcolIygTV9Z0GaeX8puRKO5PEisvWDt3qbnrdFhf1wAMkHxQ_2jYmSJ";

// Some more random base64 strings
//     "jt1vZpX9Fi6qHnRhrTb3-UuakzK0_JEW47wxeCO8Qf5IlgPsoYScDm2yNLdGAMBV";
//     "PGpN7Ws0gaFR6mvJT1UXl3bHBxtnuiyq-d9fj_wYckV2zSKIoA5rLMOeDC4ZhEQ8";
//     "3mqEds5hZkUyjD269ABplRHgI8iYzr-XOLbwF07ctou1SveV4KGQCMPNaxnTWfJ_";
//     "b9Tm5k27HuB-VyLEl13RIMwSKNGUDYQpnPhsJgavc6OiC4ofjFrAxd0_ztZqWX8e";
//     "Na3gFiQx1sS8LKyOuZrYBpjzwGEDPbomdq654RcIX_0e2C7k-WHnUhVAJlMf9Ttv";
//     "Zw84hDk-pN5uKcPy1_LdqIn0tQCJWBAROm26XSijzegxME7FHVbUaTGlorf3sYv9";
//     "fBe04QGcSkXsLud76gbxIFpOyUHajWiZmYMrEnDhtw5KqCRA8v3lPTz_o12VN9-J";
//     "RoFTY_jOZtbkai8651lp-VqzEgd4rLuDJ2WHBUv3xA9C0m7wKnsPhfMSQecGINXy";
//     "r3fi0dH_6kYyOaQ8s2eUBWucGS7PnNq9moFbTEh4C1xwMXJzIv-VZDljtRgLA5pK";
//     "vjOShxu1Cq8-JBsylNTGoiX5Kpt0cAEZr9VP2HMw3mkzFI4YL_bfRUegDWn7Qa6d";

unsigned int unRandomBase64[128];

std::string encode(uint64_t value)
{
    std::string code;
    code.reserve(11);
    uint64_t accumulate = 0, remainder = 0, position;

    do
    {
        accumulate += remainder;
        remainder = value & 0x3f;
        value >>= 6;
        position = (accumulate + remainder) & 0x3f;
        code.push_back(randomBase64[position]);
    } while (value > 0);

    return code;
}

uint64_t decode(std::string &code)
{
    uint64_t value = 0, accumulate = 0, remainder, position;

    size_t size = code.size();
    for (size_t i = 0; i < size; i++)
    {
        position = unRandomBase64[int(code[i])];
        remainder = (position + 64 - accumulate) & 0x3f;
        accumulate += remainder;
        value += remainder << (6 * i);
    }

    return value;
}

int main()
{
    // Reverse random base64 into an array.
    for (int i = 0; i < 64; i++)
    {
        unRandomBase64[int(randomBase64[i])] = i;
    }

    // Starts from 10^19.
    uint64_t step = 6553600ull;
    for (uint64_t i = 10000000000000000000ull; i < 10000000000000000000ull + step * 100ull; i += step)
    {
        for (uint64_t num = i; num < i + step; num++)
        {
            std::string code = encode(num);
            decode(code);
            // uint64_t value = decode(code);
            // std::cout << num << " -> " << code << " -> " << value << std::endl;

            // if (num != value)
            // {
            //     std::cout << "Decode Error " << num << " -> " << code << " -> " << value << std::endl;
            // }
        }

        std::cout << "Completed calculation of range [" << i << ", " << i + step << ")." << std::endl;
    }

    return 0;
}
