#include <sysdev.h>

/**
 * @brief Get the Network Card Info List Item object
 */
static NetworkCardInfo getNetworkCardInfoListItem(NetworkCardInfo* list, size_t index) {
    return list[index];
}